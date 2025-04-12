package example_server_app

import (
	"database/sql"
	"net/http"
	"rz-server/internal/app/example/api/example"
	"rz-server/internal/app/example/application/example"
	"rz-server/internal/app/example/domain/example"
	in_memory_consume_event "rz-server/internal/app/example/infra/events/in_memory_consume"
	example_sql_store "rz-server/internal/app/example/infra/store/sql/example"
	"rz-server/internal/common/interfaces"
	"rz-server/internal/common/middlewares"
)

var _ interfaces.ServerApp = (*ServerApp)(nil)

type ServerApp struct {
	server interfaces.Server
	event  <-chan interfaces.Event
	util   *interfaces.Util
	sqlDB  *sql.DB
}

func New(cmd *interfaces.CMD) *ServerApp {
	return &ServerApp{
		server: cmd.Server,
		event:  cmd.ConsumeEvent,
		util:   cmd.Util,
		sqlDB:  cmd.SqlDB,
	}
}

func (exampleApp *ServerApp) RegisterAPI() error {
	exampleApp.server.RegisterMiddlewares([]func(http.Handler) http.Handler{
		middlewares.NewLoggingMiddleware(exampleApp.util.Log),
	})

	example_store := example_sql_store.New(exampleApp.sqlDB, exampleApp.util)
	example_domain := example.NewEntity()
	exampleService := example_service.New(example_store, example_domain)

	example_api.New(
		exampleApp.server,
		exampleService,
		exampleApp.util,
	).Register()

	return nil
}

func (exampleApp *ServerApp) RegisterDomainEvent() error {
	go func() {
		example_store := example_sql_store.New(exampleApp.sqlDB, exampleApp.util)
		example_domain := example.NewEntity()

		exampleService := example_service.New(example_store, example_domain)

		in_memory_consume_event.NewInMemoryExampleEventConsumer(exampleApp.event, exampleService).StartEventLoop()

	}()

	return nil
}
