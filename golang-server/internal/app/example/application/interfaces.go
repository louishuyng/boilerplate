package application

import (
	example_commands "rz-server/internal/app/example/application/example/commands"
	example_data "rz-server/internal/app/example/application/example/data"
	"rz-server/internal/common/errors/application_error"
)

type ExampleService interface {
	CreateExample(command example_commands.CreateExampleCommand) (*example_data.ExampleData, *application_error.Error)
}

type TokenService interface {
}

type AuthService interface {
}
