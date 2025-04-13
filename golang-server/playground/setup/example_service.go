package playground_setup

import (
	"rz-server/internal/app/example/application"
	"rz-server/internal/app/example/application/example"
	"rz-server/internal/app/example/domain/example"
	sql_store "rz-server/internal/app/example/infra/store/sql"
	example_sql_store "rz-server/internal/app/example/infra/store/sql/example"
)

func MakeExampleService() application.ExampleService {
	cmd := MakeCMD()
	repository := sql_store.NewRepository(cmd.SqlDB, cmd.Util)

	example_store := example_sql_store.New(repository)
	example_domain := example.New()

	exampleService := example_service.New(example_store, example_domain)

	return exampleService
}
