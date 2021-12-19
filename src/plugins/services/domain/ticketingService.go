package domain

import (
	"github.com/dbond762/go_services_aggregator/src/plugins/services/domain/models"
)

type TicketingService interface {
	Service
	GetAllTickets() ([]models.Ticket, error)
}
