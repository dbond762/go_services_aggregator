package objects

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/dbond762/go_services_aggregator/src/plugins/services/domain/models"
)

const (
	ticketsTable = "tickets"
)

type TicketsObject interface {
	BatchAdd([]models.Ticket) error
	Clear() error
	GetListByUserID(idUser int64) ([]models.Ticket, error)
}

type ticketsObject struct {
	db *sql.DB
}

func NewTicketsObject(db *sql.DB) TicketsObject {
	return &ticketsObject{db}
}

func (object ticketsObject) BatchAdd(tickets []models.Ticket) error {
	valuesStrings := make([]string, len(tickets))
	valuesArgs := make([]interface{}, 0, len(tickets)*9)

	for i, ticket := range tickets {
		valuesStrings[i] = "(?, ?, ?, ?, ?, ?, ?, ?, ?)"
		valuesArgs = append(
			valuesArgs,
			ticket.UserServiceID,
			ticket.Name,
			ticket.Type,
			ticket.Project,
			ticket.Caption,
			ticket.Status,
			ticket.Priority,
			ticket.Assignee,
			ticket.Creator,
		)
	}

	query := fmt.Sprintf("INSERT INTO %s"+
		" (id_user_service, name, type, project, caption, status, priority, assignee, creator) "+
		"VALUES %s", ticketsTable, strings.Join(valuesStrings, ","))

	_, err := object.db.Exec(query, valuesArgs...)

	return err
}

func (object ticketsObject) Clear() error {
	query := "DELETE FROM " + ticketsTable

	_, err := object.db.Exec(query)

	return err
}

func (object ticketsObject) GetListByUserID(idUser int64) ([]models.Ticket, error) {
	query := `
		SELECT
			` + ticketsTable + `.id,
			` + ticketsTable + `.name,
			` + ticketsTable + `.type,
			` + ticketsTable + `.project,
			` + ticketsTable + `.caption,
			` + ticketsTable + `.status,
			` + ticketsTable + `.priority,
			` + ticketsTable + `.assignee,
			` + ticketsTable + `.creator
		FROM
		` + ticketsTable + `
			INNER JOIN ` + usersTable + ` ON ` + ticketsTable + `.id_user_service = ` + usersTable + `.id
		WHERE ` + usersTable + `.id_user = ?`

	rows, err := object.db.Query(query, idUser)
	if err != nil {
		return nil, err
	}

	tickets := make([]models.Ticket, 0)

	for rows.Next() {
		ticket := new(models.Ticket)

		err = rows.Scan(
			&ticket.ID,
			&ticket.Name,
			&ticket.Type,
			&ticket.Project,
			&ticket.Caption,
			&ticket.Status,
			&ticket.Priority,
			&ticket.Assignee,
			&ticket.Creator,
		)
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, *ticket)
	}

	return tickets, err
}
