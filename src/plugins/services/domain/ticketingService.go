package domain

import (
	"github.com/dbond762/go_services_aggregator/src/plugins/services/models"
)

type TicketingService interface {
	Service

	SearchAll() ([]models.Ticket, error)
}
