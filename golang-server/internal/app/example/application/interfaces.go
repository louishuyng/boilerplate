package application

import "rz-server/internal/app/example/application/commands"

type Example struct {
	ID   int
	Name string
}

type ExampleService interface {
	CreateExample(command commands.CreateExampleCommand) (*Example, error)
}
