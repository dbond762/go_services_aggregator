package settings

import (
	"database/sql"
	"github.com/dbond762/go_services_aggregator/src/libs/session"
	"github.com/dbond762/go_services_aggregator/src/theme"
	"net/http"
)

type Handler struct {
	session *session.Manager
	db      *sql.DB
	theme   *theme.Theme
}

func NewHandler(session *session.Manager, db *sql.DB, theme *theme.Theme) *Handler {
	return &Handler{
		session: session,
		db:      db,
		theme:   theme,
	}
}

func (h Handler) Settings(w http.ResponseWriter, r *http.Request) {
}
