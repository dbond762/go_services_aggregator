package main

import "github.com/dbond762/go_services_aggregator/src/theme"

func newTheme() *theme.Theme {
	const title = "Service Aggregator"

	t := theme.NewTheme(title)
	t.Init()

	t.AddMenuItem("/ticketing/", "Ticketing", 10)

	return t
}
