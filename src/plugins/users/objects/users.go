package objects

import "github.com/dbond762/go_services_aggregator/src/plugins/users/models"

type Users interface {
	GetByUsername(username string) (*models.User, error)
	Create(username string, password []byte) error
}
