package application

import (
	example_commands "rz-server/internal/app/example/application/example/commands"
	example_data "rz-server/internal/app/example/application/example/data"
)

type ExampleService interface {
	CreateExample(command example_commands.CreateExampleCommand) (*example_data.ExampleData, error)
}

type TokenService interface {
}

type AuthService interface {
}
