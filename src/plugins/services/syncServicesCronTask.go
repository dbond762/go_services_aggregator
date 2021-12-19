package services

import (
	"database/sql"
	"log"

	"github.com/dbond762/go_services_aggregator/src/plugins/services/domain"
	"github.com/dbond762/go_services_aggregator/src/plugins/services/domain/models"
	"github.com/dbond762/go_services_aggregator/src/plugins/services/objects"

	_ "github.com/dbond762/go_services_aggregator/src/plugins/services/domain/jira"
)

type SyncServicesCronTask struct {
	servicesObject objects.ServicesObject
	ticketsObject  objects.TicketsObject
}

func NewSyncServicesCronTask(db *sql.DB) *SyncServicesCronTask {
	return &SyncServicesCronTask{
		servicesObject: objects.NewServicesObject(db),
		ticketsObject:  objects.NewTicketsObject(db),
	}
}

func (task SyncServicesCronTask) Execute() {
	defer func() {
		if err := recover(); err != nil {
			log.Print(err)
		}
	}()

	services, err := task.servicesObject.GetAll()
	if err != nil {
		panic(err)
	}

	for _, service := range services {
		if err := task.executeService(service); err != nil {
			panic(err)
		}
	}
}

func (task SyncServicesCronTask) executeService(serviceModel models.Service) error {
	serviceInstance, err := domain.CreateService(serviceModel.Ident)
	if err != nil {
		return err
	}

	if err := serviceInstance.Init(serviceModel); err != nil {
		return err
	}

	switch serviceInstance.(type) {
	case domain.TicketingService:
		err = task.executeTicketingService(serviceInstance.(domain.TicketingService))
	}

	if err != nil {
		return err
	}

	serviceInstance.Finalize()

	return nil
}

func (task SyncServicesCronTask) executeTicketingService(service domain.TicketingService) error {
	tickets, err := service.GetAllTickets()
	if err != nil {
		return err
	}

	if err := task.ticketsObject.BatchAdd(tickets); err != nil {
		return err
	}

	return nil
}
