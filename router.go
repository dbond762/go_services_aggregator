package main

import (
	"net/http"

	"github.com/dbond762/go_services_aggregator/src/libs/session"
	"github.com/dbond762/go_services_aggregator/src/plugins/services"
	"github.com/dbond762/go_services_aggregator/src/plugins/users"
	"go.uber.org/dig"
)

func initRoutes(container *dig.Container) error {
	if err := container.Provide(users.NewHandler); err != nil {
		return err
	}

	if err := container.Provide(services.NewHandler); err != nil {
		return err
	}

	return container.Invoke(setupRouters)
}

func setupRouters(usersHandler *users.Handler, servicesHandler *services.Handler, session *session.Manager) {
	http.HandleFunc("/login/", usersHandler.Login)
	http.HandleFunc("/logout/", usersHandler.Logout)

	http.HandleFunc("/", servicesHandler.Index)
	http.HandleFunc("/ticketing/", users.Auth(session, servicesHandler.Ticketing))
}
