package services

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dbond762/go_services_aggregator/src/libs/session"
	"github.com/dbond762/go_services_aggregator/src/plugins/services/objects"
	"github.com/dbond762/go_services_aggregator/src/plugins/services/storage"
	usersModels "github.com/dbond762/go_services_aggregator/src/plugins/users/models"
	"github.com/dbond762/go_services_aggregator/src/theme"
)

type Handler struct {
	ticketsObject objects.Tickets
	theme         *theme.Theme
	session       *session.Manager
}

func NewHandler(theme *theme.Theme, session *session.Manager, db *sql.DB) *Handler {
	return &Handler{
		ticketsObject: storage.NewTickets(db),
		theme:         theme,
		session:       session,
	}
}

func (h Handler) Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ticketing/", http.StatusSeeOther)
}

func (h Handler) Ticketing(w http.ResponseWriter, r *http.Request) {
	sess := h.session.SessionStart(w, r)
	user := sess.Get("user").(*usersModels.User)

	paths := []string{
		"src/plugins/services/templates/ticketing.html",
	}

	tickets, err := h.ticketsObject.GetListByUserID(user.ID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	h.theme.Display(w, paths, tickets)
}
