package storage

import (
	"database/sql"

	"github.com/dbond762/go_services_aggregator/src/plugins/settings/models"
	"github.com/dbond762/go_services_aggregator/src/plugins/settings/objects"
)

const tableCredentials = "`credentials`"

type credentials struct {
	db *sql.DB
}

func NewCredentials(db *sql.DB) objects.Credentials {
	return &credentials{db}
}

func (c credentials) GetAll() (results map[int64][]models.Credential, err error) {
	query := "SELECT `id_setting`, `key`, `value` FROM " + tableCredentials

	rows, err := c.db.Query(query)
	if err != nil {
		return
	}

	results = make(map[int64][]models.Credential, 0)

	for rows.Next() {
		var (
			settingID int64
			key       string
			value     string
		)

		err = rows.Scan(&settingID, &key, &value)
		if err != nil {
			return
		}

		credential := models.Credential{
			SettingID: settingID,
			Key:       key,
			Value:     value,
		}

		results[settingID] = append(results[settingID], credential)
	}

	return
}
