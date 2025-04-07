package main

import (
	"go-server/internal/common"
	"go-server/internal/user"
)

func main() {
	server := NewServer()

	event := common.NewEventChannel()

	userApp := user.NewUserApp(server, event)

	_ = userApp.RunServer()

	go func() {
		_ = userApp.RunDomainEventLoop()
	}()
}
