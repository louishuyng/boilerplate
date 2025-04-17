package example_service

import (
	example_commands "rz-server/internal/app/example/application/example/commands"
	example_data "rz-server/internal/app/example/application/example/data"
	"rz-server/internal/common/interfaces"
)

func (u *ExampleService) CreateExample(command example_commands.CreateExampleCommand) (*example_data.ExampleData, interfaces.ApplicationError) {
	data := u.store.CreateExample("Hello")

	return &example_data.ExampleData{
		ID:   data.Id,
		Name: data.Name,
	}, nil
}
