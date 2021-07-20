package users

import (
	"net/http"

	"github.com/dbond762/go_services_aggregator/src/libs/session"
)

func Auth(session *session.Manager, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sess := session.SessionStart(w, r)
		user := sess.Get("user")
		if user == nil {
			http.Redirect(w, r, "/login/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	}
}
