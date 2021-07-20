package storage

import (
	"database/sql"

	"github.com/dbond762/go_services_aggregator/src/plugins/users/models"
	"github.com/dbond762/go_services_aggregator/src/plugins/users/objects"
	"golang.org/x/crypto/bcrypt"
)

const tableUsers = "users"

type users struct {
	db *sql.DB
}

func NewUserObject(db *sql.DB) objects.Users {
	return &users{db}
}

func (u users) GetByUsername(username string) (*models.User, error) {
	query := "SELECT * FROM " + tableUsers + " WHERE username = ?"

	row := u.db.QueryRow(query, username)

	var id int64
	var password []byte
	if err := row.Scan(&id, &username, &password); err != nil {
		return nil, err
	}

	user := models.NewUser(id, username, password)
	return user, nil
}

func (u users) Create(username string, password []byte) error {
	query := "INSERT INTO " + tableUsers + " (username, password) VALUES (?, ?)"

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if _, err := u.db.Exec(query, username, hashedPassword); err != nil {
		return err
	}

	return nil
}
