package main

import (
	"database/sql"
	"rz-server/internal/app/example"
	"rz-server/internal/common/interfaces"
	"rz-server/internal/common/message_brokers"
)

func main() {
	log := NewLog()
	env := newEnvironment()

	util := interfaces.Util{
		Log:         NewLog(),
		Logger:      log.ErrorLogger,
		Environment: env,
	}

	server := NewServer(&util)
	event := message_brokers.NewEventChannel()

	sql, err := NewSQLConnection(SQLSettings{
		Host:     env.GetEnv("database", "HOST"),
		Port:     env.GetEnv("database", "PORT"),
		User:     env.GetEnv("database", "USER"),
		Password: env.GetEnv("database", "PASSWORD"),
		DBName:   env.GetEnv("database", "DATABASE"),
	}, log)

	if err != nil {
		log.Error("Failed to connect to database", map[string]any{
			"error": err.Error(),
		})
		panic(err)
	}

	cmd := makeCMD(server, event, &util, sql)

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

func makeCMD(
	server interfaces.Server,
	event <-chan interfaces.Event,
	util *interfaces.Util,
	sqlDb *sql.DB,
) interfaces.CMD {
	return interfaces.CMD{
		Server:       server,
		ConsumeEvent: event,
		Util:         util,
		SqlDB:        sqlDb,
	}
}
