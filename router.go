package main

import (
	"log"
	"net/http"

	"github.com/dbond762/go_services_aggregator/src/libs/session"
	"github.com/dbond762/go_services_aggregator/src/plugins/services"
	"github.com/dbond762/go_services_aggregator/src/plugins/settings"
	"github.com/dbond762/go_services_aggregator/src/plugins/users"
	"go.uber.org/dig"
)

func initRoutes(container *dig.Container) {
	if err := container.Provide(users.NewHandler); err != nil {
		log.Fatal(err)
	}

	if err := container.Provide(services.NewHandler); err != nil {
		log.Fatal(err)
	}

	if err := container.Provide(settings.NewHandler); err != nil {
		log.Fatal(err)
	}

	if err := container.Invoke(setupRouters); err != nil {
		log.Fatal(err)
	}
}

func setupRouters(
	usersHandler *users.Handler,
	servicesHandler *services.Handler,
	settingsHandler *settings.Handler,
	session *session.Manager,
) {
	http.HandleFunc("/login/", usersHandler.Login)
	http.HandleFunc("/logout/", usersHandler.Logout)

	http.HandleFunc("/", servicesHandler.Index)
	http.HandleFunc("/ticketing/", users.Auth(session, servicesHandler.Ticketing))

	http.HandleFunc("/settings/", users.Auth(session, settingsHandler.Settings))
}
