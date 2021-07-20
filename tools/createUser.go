package tools

import (
	"database/sql"

	"github.com/dbond762/go_services_aggregator/src/plugins/users/storage"
)

func CreateUser(db *sql.DB, username string, password []byte) error {
	usersObject := storage.NewUserObject(db)

	return usersObject.Create(username, password)
}
