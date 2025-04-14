package example_service

import (
	example_commands "rz-server/internal/app/example/application/example/commands"
	example_data "rz-server/internal/app/example/application/example/data"
	example_errors "rz-server/internal/app/example/application/example/errors"
	"rz-server/internal/common/interfaces"
)

func (u *ExampleService) CreateExample(command example_commands.CreateExampleCommand) (*example_data.ExampleData, interfaces.ApplicationError) {
	u.store.CreateExample("Hello")
	return nil, u.errors.New(example_errors.EXAMPLE_NOT_FOUND)
}
