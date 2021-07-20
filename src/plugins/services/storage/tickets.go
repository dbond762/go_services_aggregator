package storage

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/dbond762/go_services_aggregator/src/plugins/services/models"
	"github.com/dbond762/go_services_aggregator/src/plugins/services/objects"
)

const tableTickets = "tickets"

type tickets struct {
	db *sql.DB
}

func NewTickets(db *sql.DB) objects.Tickets {
	return &tickets{db}
}

func (t tickets) BatchAdd(tickets []models.Ticket) error {
	valuesStrings := make([]string, len(tickets))
	valuesArgs := make([]interface{}, 0, len(tickets)*9)

	for i, ticket := range tickets {
		valuesStrings[i] = "(?, ?, ?, ?, ?, ?, ?, ?, ?)"
		valuesArgs = append(
			valuesArgs,
			ticket.UserID,
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
		" (id_user, name, type, project, caption, status, priority, assignee, creator) "+
		"VALUES %s", tableTickets, strings.Join(valuesStrings, ","))

	_, err := t.db.Exec(query, valuesArgs...)

	return err
}

func (t tickets) Clear() error {
	query := "DELETE FROM " + tableTickets

	_, err := t.db.Exec(query)

	return err
}
