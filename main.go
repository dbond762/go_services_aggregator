package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/dbond762/go_services_aggregator/src/libs/session/providers/memory"
	// "github.com/dbond762/go_services_aggregator/tools"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	container, err := newContainer()
	if err != nil {
		log.Fatal(err)
	}

	initRoutes(container)
	// initCronServices(container)

	// tools.CreateUser(db, "admin", []byte("123456"))

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
