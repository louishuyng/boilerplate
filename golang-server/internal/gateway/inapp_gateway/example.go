package inapp_gateway_example

import (
	"database/sql"
	example_application "rz-server/internal/app/example/application"
	example_service "rz-server/internal/app/example/application/example"
	"rz-server/internal/app/example/domain/example"
	example_sql_store "rz-server/internal/app/example/infra/store/sql/example"
	"rz-server/internal/common/interfaces"
	"rz-server/internal/gateway"
)

var _ gateway.ExampleGateway = (*ExampleGateway)(nil)

type ExampleGateway struct {
	exampleService example_application.ExampleService
}

func New(db *sql.DB, util *interfaces.Util) *ExampleGateway {
	example_store := example_sql_store.New(db, util)
	example_domain := example.New()

	exampleService := example_service.New(example_store, example_domain)

	return &ExampleGateway{
		exampleService: exampleService,
	}
}

func (e *ExampleGateway) ExampleGet(id string) (string, error) {
	return "", nil
}
