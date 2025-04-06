package main

import (
	"go-server/internal/common"
	"go-server/internal/user"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	event := common.NewEventChannel()

	user.NewUserApp(router, event).Run()

	log.Fatal(http.ListenAndServe(":5001", router))
}
