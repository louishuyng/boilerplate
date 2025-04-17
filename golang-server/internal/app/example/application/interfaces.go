package application

import (
	example_commands "rz-server/internal/app/example/application/example/commands"
	example_data "rz-server/internal/app/example/application/example/data"
	"rz-server/internal/common/interfaces"
)

type ExampleService interface {
	CreateExample(command example_commands.CreateExampleCommand) (*example_data.ExampleData, interfaces.ApplicationError)
	ValidateToken(token string) bool
}
