package domain

import "github.com/dbond762/go_services_aggregator/src/plugins/services/domain/models"

type Service interface {
	Init(service models.Service) error
	Finalize()

	CredentialsKeys() []string
}
