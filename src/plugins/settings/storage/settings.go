package storage

import (
	"database/sql"
	"fmt"

	"github.com/dbond762/go_services_aggregator/src/plugins/settings/models"
	"github.com/dbond762/go_services_aggregator/src/plugins/settings/objects"
)

const (
	tableSettings = "settings"
	tableServices = "services"
)

type settings struct {
	db *sql.DB
}

func NewSettings(db *sql.DB) objects.Settings {
	return &settings{db}
}

func (s settings) GetAll() (results []models.Setting, err error) {
	query := fmt.Sprintf(
		"SELECT %[1]s.id, %[1]s.id_user, %[1]s.service, %[2]s.type "+
			"FROM %[1]s "+
			"INNER JOIN %[2]s ON %[1]s.service = %[2]s.ident",
		tableSettings,
		tableServices,
	)

	rows, err := s.db.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var (
			id          int64
			userID      int64
			service     string
			serviceType string
		)

		err = rows.Scan(&id, &userID, &service, &serviceType)
		if err != nil {
			return
		}

		setting := models.Setting{
			ID:          id,
			UserID:      userID,
			Service:     service,
			ServiceType: serviceType,
		}

		results = append(results, setting)
	}

	return
}
