package example_server_app

import (
	"database/sql"
	"net/http"
	"rz-server/internal/app/example/api"
	"rz-server/internal/app/example/application"
	"rz-server/internal/app/example/application/service/example"
	"rz-server/internal/app/example/domain/example"
	in_memory_consume_event "rz-server/internal/app/example/infra/events/in_memory_consume"
	sql_store "rz-server/internal/app/example/infra/store/sql"
	"rz-server/internal/common/interfaces"
	"rz-server/internal/common/middlewares"
)

var _ interfaces.ServerApp = (*ServerApp)(nil)

type ServerApp struct {
	server  interfaces.Server
	service application.ExampleService
	event   <-chan interfaces.Event
	util    *interfaces.Util
	sqlDB   *sql.DB
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
	store := sql_store.New(exampleApp.sqlDB, exampleApp.util.Log)
	domain := example.New()

	exampleApp.server.RegisterMiddlewares([]func(http.Handler) http.Handler{
		middlewares.NewLoggingMiddleware(exampleApp.util.Log),
	})

	exampleService := example_service.New(store, domain)

	example_api.New(
		exampleApp.server,
		exampleService,
		exampleApp.util,
	).Register()

	return nil
}

func (exampleApp *ServerApp) RegisterDomainEvent() error {
	go func() {
		in_memory_consume_event.NewInMemoryExampleEventConsumer(exampleApp.event, exampleApp.service).StartEventLoop()

	}()

	return nil
}
