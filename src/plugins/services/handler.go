package services

import (
	"net/http"

	"github.com/dbond762/go_services_aggregator/src/libs/session"
	"github.com/dbond762/go_services_aggregator/src/theme"
)

type Handler struct {
	theme   *theme.Theme
	session *session.Manager
}

func NewHandler(theme *theme.Theme, session *session.Manager) *Handler {
	return &Handler{
		theme:   theme,
		session: session,
	}
}

func (h Handler) Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ticketing/", http.StatusSeeOther)
}

func (h Handler) Ticketing(w http.ResponseWriter, r *http.Request) {
	paths := []string{
		"src/plugins/services/templates/ticketing.html",
	}

	data := map[string]string{
		"Username": "Vasiliy",
	}

	h.theme.Display(w, paths, data)
}
