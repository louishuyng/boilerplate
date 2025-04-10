package example

import domain "rz-server/internal/app/example/domain"

type ExampleEntity struct {
	Id   int
	Name string
}

type Example struct {
}

var _ domain.Example = (*Example)(nil)

func New() *Example {
	return &Example{}
}

func (d *Example) GetID() int {
	return 0
}
