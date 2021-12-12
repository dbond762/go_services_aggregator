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

func (t tickets) GetListByUserID(userID int64) (tickets []models.Ticket, err error) {
	query := fmt.Sprintf(
		`SELECT id,
			    id_user,
				name,
				type,
				project,
				caption,
				status,
				priority,
				assignee,
				creator
		FROM %s
		WHERE id_user = ?`,
		tableTickets,
	)

	rows, err := t.db.Query(query, userID)
	if err != nil {
		return
	}

	for rows.Next() {
		var (
			id         int64
			userID     int64
			name       string
			ticketType string
			project    string
			caption    string
			status     string
			priority   string
			assignee   string
			creator    string
		)

		err = rows.Scan(&id, &userID, &name, &ticketType, &project, &caption, &status, &priority, &assignee, &creator)
		if err != nil {
			return
		}

		ticket := models.Ticket{
			ID:       id,
			UserID:   userID,
			Name:     name,
			Type:     ticketType,
			Project:  project,
			Caption:  caption,
			Status:   status,
			Priority: priority,
			Assignee: assignee,
			Creator:  creator,
		}

		tickets = append(tickets, ticket)
	}

	return
}
