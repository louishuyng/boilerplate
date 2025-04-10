package main

import (
	"database/sql"
	"fmt"
	"rz-server/internal/common/interfaces"

	_ "github.com/lib/pq"
)

type SQLSettings struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewSQLConnection(settings SQLSettings, log interfaces.LogUtil) (*sql.DB, error) {
	log.Info("Connecting to database")

	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		settings.User,
		settings.Password,
		settings.Host,
		settings.Port,
		settings.DBName,
	)

	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		log.Error("db connection failure: %v", map[string]interface{}{
			"error": err.Error(),
		})
	}

	// test db connection
	err = db.Ping()
	if err != nil {
		log.Error("db ping failure: %v", map[string]interface{}{
			"error": err.Error(),
		})
	}

	log.Info("db connection successful")

	return db, nil
}
