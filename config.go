package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Address string
	DSN     string
}

func getConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Address: os.Getenv("ADDR"),
		DSN:     os.Getenv("DB_DSN"),
	}
}
