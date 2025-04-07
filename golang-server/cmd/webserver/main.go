package main

import (
	"go-server/internal/common"
	"go-server/internal/user"
)

func main() {
	server := NewServer()

	event := common.NewEventChannel()

	userApp := user.NewApp(server, event)

	_ = userApp.RegisterAPI()
	_ = userApp.RegisterDomainEvent()

	server.Start()
}
