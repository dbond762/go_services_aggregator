package services

import (
	"database/sql"
	"log"

	"github.com/dbond762/go_services_aggregator/src/plugins/services/domain"
	"github.com/dbond762/go_services_aggregator/src/plugins/services/domain/jira"
	"github.com/dbond762/go_services_aggregator/src/plugins/services/objects"
	"github.com/dbond762/go_services_aggregator/src/plugins/services/storage"
	"github.com/dbond762/go_services_aggregator/src/plugins/settings/models"
	settingsObject "github.com/dbond762/go_services_aggregator/src/plugins/settings/objects"
)

type CronService struct {
	settingsObject    settingsObject.Settings
	credentialsObject settingsObject.Credentials
	ticketsObject     objects.Tickets
}

func NewCronService(
	db *sql.DB,
	settingsObject settingsObject.Settings,
	credentialsObject settingsObject.Credentials,
) *CronService {
	return &CronService{
		settingsObject:    settingsObject,
		credentialsObject: credentialsObject,
		ticketsObject:     storage.NewTickets(db),
	}
}

func (s CronService) ExecuteService() {
	defer func() {
		if err := recover(); err != nil {
			log.Print(err)
		}
	}()

	settings, err := s.settingsObject.GetAll()
	if err != nil {
		panic(err)
	}

	credentials, err := s.credentialsObject.GetAll()
	if err != nil {
		panic(err)
	}

	if err := s.ticketsObject.Clear(); err != nil {
		panic(err)
	}

	for _, setting := range settings {
		settingCredentials := s.transformSettingCredentials(credentials[setting.ID])

		var err error

		switch setting.Service {
		case "Jira":
			err = s.executeJiraService(setting, settingCredentials)
		}

		if err != nil {
			panic(err)
		}
	}
}

func (s CronService) executeJiraService(setting models.Setting, credentials map[string]string) error {
	service := new(jira.Service)

	if err := service.Init(setting.UserID, credentials); err != nil {
		return err
	}

	if err := s.executeTicketingService(service); err != nil {
		return err
	}

	service.Finalize()

	return nil
}

func (s CronService) transformSettingCredentials(credentials []models.Credential) map[string]string {
	result := make(map[string]string, len(credentials))

	for _, credential := range credentials {
		result[credential.Key] = credential.Value
	}

	return result
}

func (s CronService) executeTicketingService(service domain.TicketingService) error {
	tickets, err := service.SearchAll()
	if err != nil {
		return err
	}

	if err := s.ticketsObject.BatchAdd(tickets); err != nil {
		return err
	}

	return nil
}
