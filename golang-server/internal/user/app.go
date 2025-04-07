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

var _ common.App = (*UserApp)(nil)

type UserApp struct {
	server  common.BasicServer
	service application.UserService
	event   <-chan common.Event
}

func NewUserApp(server common.BasicServer, event <-chan common.Event) *UserApp {
	store := store.NewPostgresStore()
	domain := domain.NewUserDomain()

	service := service.NewUserService(store, domain)

	return &UserApp{
		server:  server,
		service: service,
		event:   event,
	}
}

func (userApp *UserApp) RegisterAPI() error {
	api.NewUserApi(userApp.server, userApp.service)

	return nil
}

func (userApp *UserApp) RegisterDomainEvent() error {
	go func() {
		events.NewInMemoryUserEventConsumer(userApp.event, userApp.service).StartEventLoop()

	}()

	return nil
}
