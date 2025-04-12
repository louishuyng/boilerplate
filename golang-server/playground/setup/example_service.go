package playground_setup

import (
	"rz-server/internal/app/example/application"
	"rz-server/internal/app/example/application/example"
	"rz-server/internal/app/example/domain/example"
	example_sql_store "rz-server/internal/app/example/infra/store/sql/example"
)

func MakeExampleService() application.ExampleService {
	cmd := MakeCMD()

	example_store := example_sql_store.New(cmd.SqlDB, cmd.Util)
	example_domain := example.New()

	exampleService := example_service.New(example_store, example_domain)

	return exampleService
}
