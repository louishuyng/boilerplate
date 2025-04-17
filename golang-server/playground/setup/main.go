package playground_setup

import (
	"rz-server/internal/common/interfaces"
	"rz-server/internal/common/message_brokers"
)

func MakeCMD() interfaces.CMD {
	log := NewLog()
	env := NewEnvironment()

	util := interfaces.Util{
		Log:         log,
		Logger:      log.ErrorLogger,
		Environment: env,
	}

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

	return interfaces.CMD{
		ConsumeEvent: event,
		Util:         &util,
		SqlDB:        sql,
	}
}
