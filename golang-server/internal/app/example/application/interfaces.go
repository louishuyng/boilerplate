package application

type Example struct {
	ID   int
	Name string
}

type ExampleService interface {
	CreateExample(createExampleCommand CreateExampleCommand) (*Example, error)
}
