package main

import (
	"log"

	"github.com/dbond762/go_services_aggregator/src/plugins/services"
	settingsStorage "github.com/dbond762/go_services_aggregator/src/plugins/settings/storage"
	"github.com/robfig/cron/v3"
	"go.uber.org/dig"
)

func initCronServices(container *dig.Container) {
	if err := container.Provide(settingsStorage.NewSettings); err != nil {
		log.Fatal(err)
	}

	if err := container.Provide(settingsStorage.NewCredentials); err != nil {
		log.Fatal(err)
	}

	if err := container.Provide(services.NewCronService); err != nil {
		log.Fatal(err)
	}

	if err := container.Invoke(setupCronJobs); err != nil {
		log.Fatal(err)
	}
}

func setupCronJobs(servicesCron *services.CronService) {
	c := cron.New()

	servicesCron.ExecuteService()

	_, err := c.AddFunc("@daily", servicesCron.ExecuteService)
	if err != nil {
		log.Fatal(err)
	}

	c.Start()
}
