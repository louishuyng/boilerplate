package example_server_app

import (
	"database/sql"
	"net/http"
	auth_api "rz-server/internal/app/user/api/auth"
	auth_service "rz-server/internal/app/user/application/auth"
	"rz-server/internal/app/user/domain/auth"
	sql_store "rz-server/internal/app/user/infra/store/sql"
	auth_sql_store "rz-server/internal/app/user/infra/store/sql/auth"
	user_sql_store "rz-server/internal/app/user/infra/store/sql/user"
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

func (userApp *ServerApp) RegisterAPI() error {
	repository := sql_store.NewRepository(userApp.sqlDB, userApp.util)

	userApp.server.RegisterMiddlewares([]func(http.Handler) http.Handler{
		middlewares.NewLoggingMiddleware(userApp.util.Log),
	})

	auth_store := auth_sql_store.New(repository)
	user_store := user_sql_store.New(repository)

	auth_domain := auth.New()
	authService := auth_service.New(auth_store, user_store, auth_domain)

	auth_api.New(
		userApp.server,
		authService,
		userApp.util,
	).Register()

	return nil
}

func (exampleApp *ServerApp) RegisterDomainEvent() error {
	return nil
}
