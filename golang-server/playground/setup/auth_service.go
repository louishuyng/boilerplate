package playground_setup

import (
	"rz-server/internal/app/user/application"
	auth_service "rz-server/internal/app/user/application/auth"
	"rz-server/internal/app/user/domain/auth"
	sql_store "rz-server/internal/app/user/infra/store/sql"
	auth_sql_store "rz-server/internal/app/user/infra/store/sql/auth"
	user_sql_store "rz-server/internal/app/user/infra/store/sql/user"
)

func MakeAuthService() application.AuthService {
	cmd := MakeCMD()

	repository := sql_store.NewRepository(cmd.SqlDB, cmd.Util)

	auth_store := auth_sql_store.New(repository)
	user_store := user_sql_store.New(repository)
	auth_domain := auth.New()

	authService := auth_service.New(auth_store, user_store, auth_domain)

	return authService
}
