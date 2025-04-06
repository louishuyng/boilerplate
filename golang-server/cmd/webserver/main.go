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

	userApp := user.NewUserApp(router, event)

	_ = userApp.RunServer()
	_ = userApp.RunDomainEventLoop()

	log.Fatal(http.ListenAndServe(":5001", router))
}
