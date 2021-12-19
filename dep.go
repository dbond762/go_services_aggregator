package main

import (
	"go.uber.org/dig"
)

func newContainer() (*dig.Container, error) {
	container := dig.New()

	if err := container.Provide(getConfig); err != nil {
		return nil, err
	}

	if err := container.Provide(getDBConnection); err != nil {
		return nil, err
	}

	if err := container.Provide(newSession); err != nil {
		return nil, err
	}

	if err := container.Provide(newTheme); err != nil {
		return nil, err
	}

	return container, nil
}
