package objects

import (
	"github.com/dbond762/go_services_aggregator/src/plugins/settings/models"
)

type Settings interface {
	GetAll() ([]models.Setting, error)
}
