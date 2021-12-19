package objects

import (
	"database/sql"
	"strings"

	"github.com/dbond762/go_services_aggregator/src/plugins/services/domain/models"
)

const (
	servicesTable    = "services"
	credentialsTable = "credentials"
	usersTable       = "users_services"
)

type ServicesObject interface {
	GetAll() ([]models.Service, error)
}

type servicesObject struct {
	db *sql.DB
}

func NewServicesObject(db *sql.DB) ServicesObject {
	return &servicesObject{db}
}

func (object servicesObject) GetAll() ([]models.Service, error) {
	query := "" +
		"SELECT" +
		"    " + servicesTable + ".ident,\n" +
		"    " + servicesTable + ".`type`,\n" +
		"    " + usersTable + ".id_user,\n" +
		"    " + usersTable + ".id AS id_user_service,\n" +
		"    GROUP_CONCAT(" + credentialsTable + ".`key`, ':', " + credentialsTable + ".value SEPARATOR '|') AS credentials\n" +
		"FROM\n" +
		"    " + usersTable + "\n" +
		"    INNER JOIN " + servicesTable + " ON " + usersTable + ".id_service = " + servicesTable + ".id\n" +
		"    LEFT JOIN " + credentialsTable + " ON " + credentialsTable + ".id_user_service = " + usersTable + ".id\n" +
		"GROUP BY " + usersTable + ".id"

	rows, err := object.db.Query(query)
	if err != nil {
		return nil, err
	}

	services := make([]models.Service, 0)

	for rows.Next() {
		service := new(models.Service)
		var credentials string

		err = rows.Scan(&service.Ident, &service.Type, &service.UserID, &service.UserServiceID, &credentials)
		if err != nil {
			return nil, err
		}

		service.Credentials = make(map[string]string)

		entries := strings.Split(credentials, "|")
		for _, entry := range entries {
			x := strings.Split(entry, ":")
			service.Credentials[x[0]] = x[1]
		}

		services = append(services, *service)
	}

	return services, nil
}
