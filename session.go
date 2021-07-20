package main

import "github.com/dbond762/go_services_aggregator/src/libs/session"

func newSession() (*session.Manager, error) {
	const (
		cookieName  = "idSession"
		maxLifeTime = 3600
	)

	globalSession, err := session.NewManager("memory", cookieName, maxLifeTime)
	if err != nil {
		return nil, err
	}

	go globalSession.GC()

	return globalSession, nil
}
