package inapp_gateway_example

import (
	"database/sql"
	example_application "rz-server/internal/app/example/application"
	example_service "rz-server/internal/app/example/application/service/example"
	"rz-server/internal/app/example/domain/example"
	sql_store "rz-server/internal/app/example/infra/store/sql"
	"rz-server/internal/common/interfaces"
	"rz-server/internal/gateway"
)

var _ gateway.ExampleGateway = (*ExampleGateway)(nil)

type ExampleGateway struct {
	exampleService example_application.ExampleService
}

func New(db *sql.DB, log interfaces.LogUtil) *ExampleGateway {
	store := sql_store.New(db, log)
	domain := example.New()

	exampleService := example_service.New(store, domain)

	return &ExampleGateway{
		exampleService: exampleService,
	}
}

func (e *ExampleGateway) ExampleGet(id string) (string, error) {
	return "", nil
}
