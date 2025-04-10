package example

import (
	"net/http"
	"rz-server/internal/app/example/api"
	"rz-server/internal/app/example/application"
	"rz-server/internal/app/example/application/service"
	"rz-server/internal/app/example/domain"
	"rz-server/internal/app/example/infra/events"
	"rz-server/internal/app/example/infra/store/sql"
	"rz-server/internal/common/interfaces"
	"rz-server/internal/common/middlewares"
)

var _ interfaces.App = (*App)(nil)

type App struct {
	server  interfaces.Server
	service application.ExampleService
	event   <-chan interfaces.Event
	util    *interfaces.Util
}

func NewServerApp(cmd *interfaces.CMD) *App {
	store := sql.NewSqlStore(cmd.SqlDB, cmd.Util.Log)
	domain := domain.NewExampleDomain()

	service := service.NewExampleService(store, domain)

	return &App{
		server:  cmd.Server,
		service: service,
		event:   cmd.ConsumeEvent,
		util:    cmd.Util,
	}
}

func (exampleApp *App) RegisterAPI() error {
	exampleApp.server.RegisterMiddlewares([]func(http.Handler) http.Handler{
		middlewares.NewLoggingMiddleware(exampleApp.util.Log),
	})

	api.RegisterAPI(exampleApp.server, exampleApp.service, exampleApp.util)

	return nil
}

func (exampleApp *App) RegisterDomainEvent() error {
	go func() {
		events.NewInMemoryExampleEventConsumer(exampleApp.event, exampleApp.service).StartEventLoop()

	}()

	return nil
}
