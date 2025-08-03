package main

import (
	"log"

	"golang-mongo/config"
	"golang-mongo/internal/adapters/db"
	"golang-mongo/internal/adapters/http"
	"golang-mongo/internal/application/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatal(err)
	}

	apiApp := api.NewApplication(dbAdapter)
	httpAdapter := http.NewAdapter(apiApp)

	server := http.NewServer(httpAdapter)

	log.Fatal(server.Run(config.GetApplicationPort()))
}
