package main

import (
	"database/sql"
	"time"
)

var db *sql.DB

func dbConnection(conf *Config) (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	var err error

	db, err = sql.Open("mysql", conf.DSN)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
