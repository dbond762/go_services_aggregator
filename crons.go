package main

import (
	"log"

	"github.com/dbond762/go_services_aggregator/src/plugins/services"
	"github.com/robfig/cron/v3"
	"go.uber.org/dig"
)

func initCronServices(container *dig.Container) error {
	if err := container.Provide(services.NewSyncServicesCronTask); err != nil {
		return err
	}

	return container.Invoke(setupCronJobs)
}

func setupCronJobs(syncServicesTask *services.SyncServicesCronTask) {
	c := cron.New()

	// syncServicesTask.Execute()

	_, err := c.AddFunc("@daily", syncServicesTask.Execute)
	if err != nil {
		log.Fatal(err)
	}

	c.Start()
}
