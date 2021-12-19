package jira

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/dbond762/go_services_aggregator/src/plugins/services/domain"
	"github.com/dbond762/go_services_aggregator/src/plugins/services/domain/models"
)

const (
	domainCredentialKey = "domain"
	emailCredentialKey  = "email"
	tokenCredentialKey  = "token"
)

func init() {
	domain.RegisterService("jira", func() (domain.Service, error) {
		return new(jiraService), nil
	})
}

type jiraService struct {
	userServiceID int64

	domain string
	email  string
	token  string

	client *http.Client
}

func (s *jiraService) Init(service models.Service) error {
	s.userServiceID = service.UserServiceID

	availableKeys := s.CredentialsKeys()
	for _, key := range availableKeys {
		if _, ok := service.Credentials[key]; !ok {
			return fmt.Errorf("not found %s credential", key)
		}
	}

	s.domain = service.Credentials[domainCredentialKey]
	s.email = service.Credentials[emailCredentialKey]
	s.token = service.Credentials[tokenCredentialKey]

	s.client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}

	return nil
}

func (s jiraService) Finalize() {
}

func (s jiraService) CredentialsKeys() []string {
	return []string{
		domainCredentialKey,
		emailCredentialKey,
		tokenCredentialKey,
	}
}

func (s jiraService) GetAllTickets() ([]models.Ticket, error) {
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
				UserServiceID: s.userServiceID,
				Name:          issue.Key,
				Type:          issue.Fields.IssueType.Name,
				Project:       issue.Fields.Project.Name,
				Caption:       issue.Fields.Summary,
				Status:        issue.Fields.Status.Name,
				Priority:      issue.Fields.Priority.Name,
				Assignee:      issue.Fields.Assignee.DisplayName,
				Creator:       issue.Fields.Reporter.DisplayName,
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

func (s jiraService) search(startAt int) (result Search, err error) {
	req, err := http.NewRequest(http.MethodGet, s.getUrl(startAt), nil)
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

func (s jiraService) getUrl(startAt int) string {
	params := url.Values{}
	params.Set("startAt", strconv.Itoa(startAt))

	return fmt.Sprintf("https://%s.atlassian.net/rest/api/3/search/?%s", s.domain, params.Encode())
}
