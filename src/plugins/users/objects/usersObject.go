package objects

import (
	"database/sql"

	"github.com/dbond762/go_services_aggregator/src/plugins/users/models"
	"golang.org/x/crypto/bcrypt"
)

const (
	usersTable = "users"
)

type UsersObject interface {
	GetByUsername(username string) (*models.User, error)
	Create(username string, password []byte) error
}

type usersObject struct {
	db *sql.DB
}

func NewUserObject(db *sql.DB) UsersObject {
	return &usersObject{db}
}

func (object usersObject) GetByUsername(username string) (*models.User, error) {
	query := "SELECT * FROM " + usersTable + " WHERE username = ?"

	row := object.db.QueryRow(query, username)

	var id int64
	var password []byte
	if err := row.Scan(&id, &username, &password); err != nil {
		return nil, err
	}

	user := models.NewUser(id, username, password)
	return user, nil
}

func (object usersObject) Create(username string, password []byte) error {
	query := "INSERT INTO " + usersTable + " (username, password) VALUES (?, ?)"

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if _, err := object.db.Exec(query, username, hashedPassword); err != nil {
		return err
	}

	return nil
}
