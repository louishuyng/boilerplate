package main

import (
	"rz-server/internal/app/example"
	"rz-server/internal/common/interfaces"
	"rz-server/internal/common/message_brokers"
)

func main() {
	log := NewLog()

	util := interfaces.Util{
		Log:    NewLog(),
		Logger: log.ErrorLogger,
	}

	server := NewServer(&util)
	event := message_brokers.NewEventChannel()

	cmd := makeCMD(server, event, &util)

	RegisterApp(example.NewServerApp(&cmd))

	go func() {
		server.Start()
	}()

	server.WaitForShutdown()
}

func RegisterApp(app interfaces.App) {
	_ = app.RegisterAPI()
	_ = app.RegisterDomainEvent()
}

func makeCMD(server interfaces.Server, event <-chan interfaces.Event, util *interfaces.Util) interfaces.CMD {
	return interfaces.CMD{
		Server:       server,
		ConsumeEvent: event,
		Util:         util,
	}
}
