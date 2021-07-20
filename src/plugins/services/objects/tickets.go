package objects

import "github.com/dbond762/go_services_aggregator/src/plugins/services/models"

type Tickets interface {
	BatchAdd([]models.Ticket) error
	Clear() error
}
