package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/dbond762/go_services_aggregator/src/libs/session/providers/memory"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	container, err := newContainer()
	if err != nil {
		log.Fatal(err)
	}

	if err := initRoutes(container); err != nil {
		log.Fatal(err)
	}

	if err := initCronServices(container); err != nil {
		log.Fatal(err)
	}

	if err := container.Invoke(runServer); err != nil {
		log.Fatal(err)
	}
}

func runServer(config *Config) {
	fmt.Printf("Server runing on %s\n", config.Address)
	if err := http.ListenAndServe(config.Address, nil); err != nil {
		log.Fatal(err)
	}
}
