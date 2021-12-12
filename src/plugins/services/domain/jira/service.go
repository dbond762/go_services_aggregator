package jira

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/dbond762/go_services_aggregator/src/plugins/services/domain/jira/response"
	"github.com/dbond762/go_services_aggregator/src/plugins/services/models"
)

const (
	domainCredentialKey = "domain"
	emailCredentialKey  = "email"
	tokenCredentialKey  = "token"
)

type Service struct {
	userID int64

	domain string
	email  string
	token  string

	client *http.Client
}

func (s *Service) Init(userID int64, credentials map[string]string) error {
	s.userID = userID

	availableKeys := s.CredentialsKeys()
	for _, key := range availableKeys {
		if _, ok := credentials[key]; !ok {
			return fmt.Errorf("not found %s credential", key)
		}
	}

	s.domain = credentials[domainCredentialKey]
	s.email = credentials[emailCredentialKey]
	s.token = credentials[tokenCredentialKey]

	s.client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}

	return nil
}

func (s Service) Finalize() {
}

func (s Service) CredentialsKeys() []string {
	return []string{
		domainCredentialKey,
		emailCredentialKey,
		tokenCredentialKey,
	}
}

func (s Service) Ident() string {
	return "Jira"
}

func (s Service) SearchAll() ([]models.Ticket, error) {
	tickets := make([]models.Ticket, 0)

	startAt := 0

	for {
		res, err := s.search(startAt)
		if err != nil {
			return nil, err
		}

		resultTickets := make([]models.Ticket, res.MaxResults)

		for idx, issue := range res.Issues {
			ticket := models.Ticket{
				UserID:   s.userID,
				Name:     issue.Key,
				Type:     issue.Fields.IssueType.Name,
				Project:  issue.Fields.Project.Name,
				Caption:  issue.Fields.Summary,
				Status:   issue.Fields.Status.Name,
				Priority: issue.Fields.Priority.Name,
				Assignee: issue.Fields.Assignee.DisplayName,
				Creator:  issue.Fields.Reporter.DisplayName,
			}

			resultTickets[idx] = ticket
		}

		tickets = append(tickets, resultTickets...)

		startAt += res.MaxResults

		if startAt+res.MaxResults >= res.Total {
			break
		}
	}

	return tickets, nil
}

func (s Service) search(startAt int) (result response.Search, err error) {
	req, err := http.NewRequest(http.MethodGet, s.url(startAt), nil)
	if err != nil {
		return
	}

	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(s.email, s.token)

	res, err := s.client.Do(req)
	if err != nil {
		return
	}

	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Print(err)
		}
	}()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&result)

	return
}

func (s Service) url(startAt int) string {
	params := url.Values{}
	params.Set("startAt", strconv.Itoa(startAt))

	return fmt.Sprintf("https://%s.atlassian.net/rest/api/3/search/?%s", s.domain, params.Encode())
}
