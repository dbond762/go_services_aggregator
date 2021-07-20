package main

import "flag"

type Config struct {
	Address string
	DSN     string
}

func getConfig() *Config {
	var (
		addr = flag.String("addr", ":8080", "")
		dsn  = flag.String("dsn", "root:123456@/devs_services_aggregator", "")
	)

	return &Config{
		Address: *addr,
		DSN:     *dsn,
	}
}
