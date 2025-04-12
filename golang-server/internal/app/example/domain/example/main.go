package example

import (
	domain "rz-server/internal/app/example/domain"

	"github.com/google/uuid"
)

type ExampleEntity struct {
	Id   uuid.UUID
	Name string
}

var _ domain.Example = (*ExampleEntity)(nil)

func NewEntity() *ExampleEntity {
	return &ExampleEntity{}
}

func (e *ExampleEntity) Create(
	Id uuid.UUID,
	Name string,
) ExampleEntity {
	return ExampleEntity{
		Id:   Id,
		Name: Name,
	}
}

func (d *ExampleEntity) GetID() int {
	return 0
}
