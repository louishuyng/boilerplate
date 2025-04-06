package user

import (
	"go-server/internal/common"
	"go-server/internal/user/api"
	"go-server/internal/user/application/service"
	"go-server/internal/user/domain"
	"go-server/internal/user/infra/events"
	"go-server/internal/user/infra/store"
	"net/http"
)

type UserApp struct {
	router  *http.ServeMux
	service *service.UserService
	event   <-chan common.Event
}

func NewUserApp(router *http.ServeMux, event <-chan common.Event) *UserApp {
	store := store.NewPostgresStore()
	domain := domain.NewUserDomain()

	service := service.NewUserService(store, domain)

	return &UserApp{
		router:  router,
		service: service,
		event:   event,
	}
}

func (userApp *UserApp) Run() error {
	userApp.registerUserServer()
	userApp.runDomainEventLoop()

	return nil
}

func (userApp *UserApp) registerUserServer() error {
	api.NewUserApi(userApp.router, userApp.service)

	return nil
}

func (userApp *UserApp) runDomainEventLoop() {
	events.NewInMemoryUserEventConsumer(userApp.event, userApp.service).StartEventLoop()
}
