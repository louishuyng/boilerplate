package user

import (
	"go-server/internal/common"
	"go-server/internal/user/api"
	"go-server/internal/user/application"
	"go-server/internal/user/application/service"
	"go-server/internal/user/domain"
	"go-server/internal/user/infra/events"
	"go-server/internal/user/infra/store"
)

var _ common.App = (*App)(nil)

type App struct {
	server  common.BasicServer
	service application.UserService
	event   <-chan common.Event
}

func NewApp(server common.BasicServer, event <-chan common.Event) *App {
	store := store.NewPostgresStore()
	domain := domain.NewUserDomain()

	service := service.NewUserService(store, domain)

	return &App{
		server:  server,
		service: service,
		event:   event,
	}
}

func (userApp *App) RegisterAPI() error {
	api.NewUserApi(userApp.server, userApp.service)

	return nil
}

func (userApp *App) RegisterDomainEvent() error {
	go func() {
		events.NewInMemoryUserEventConsumer(userApp.event, userApp.service).StartEventLoop()

	}()

	return nil
}
