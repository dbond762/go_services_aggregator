package objects

import "github.com/dbond762/go_services_aggregator/src/plugins/settings/models"

type Credentials interface {
	GetAll() (map[int64][]models.Credential, error)
}
