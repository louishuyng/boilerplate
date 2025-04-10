package example_service

import (
	"rz-server/internal/app/example/application"
	"rz-server/internal/app/example/application/commands"
)

func (u *ExampleService) CreateExample(command commands.CreateExampleCommand) (*application.Example, error) {
	return nil, nil
}
