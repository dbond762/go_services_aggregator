package users

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/dbond762/go_services_aggregator/src/libs/messages"
	"github.com/dbond762/go_services_aggregator/src/libs/session"
	"github.com/dbond762/go_services_aggregator/src/plugins/users/objects"
)

type Handler struct {
	userObject objects.UsersObject
	session    *session.Manager
}

func NewHandler(db *sql.DB, session *session.Manager) *Handler {
	return &Handler{
		userObject: objects.NewUserObject(db),
		session:    session,
	}
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	sess := h.session.SessionStart(w, r)
	user := sess.Get("user")
	if user != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	switch r.Method {
	case http.MethodGet:
		message := messages.Flush(w, r)
		data := loginData{
			Message: message,
		}

		paths := []string{
			"src/plugins/users/templates/login.html",
		}

		// TODO: Add recover
		t := template.Must(template.New("login.html").ParseFiles(paths...))

		if err := t.Execute(w, data); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Print(err)
			return
		}
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Print(err)
			return
		}

		// TODO: Add form verifying
		username := r.Form["username"][0]
		password := r.Form["password"][0]

		user, err := h.userObject.GetByUsername(username)
		if err != nil {
			messages.Set(w, "Invalid Login or password")
			http.Redirect(w, r, "/login/", http.StatusSeeOther)
			return
		}

		if err := user.VerifyPassword([]byte(password)); err != nil {
			messages.Set(w, "Invalid Login or password")
			http.Redirect(w, r, "/login/", http.StatusSeeOther)
			return
		}

		if err := sess.Set("user", user); err != nil {
			log.Print(err)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (h Handler) Logout(w http.ResponseWriter, r *http.Request) {
	h.session.SessionDestroy(w, r)

	http.Redirect(w, r, "/login/", http.StatusSeeOther)
}
