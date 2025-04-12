package example

import (
	domain "rz-server/internal/app/example/domain"

	"github.com/google/uuid"
)

type Entity struct {
	Id   uuid.UUID
	Name string
}

var _ domain.Example = (*Entity)(nil)

func New() *Entity {
	return &Entity{}
}

func (e *Entity) Create(
	Id uuid.UUID,
	Name string,
) Entity {
	return Entity{
		Id:   Id,
		Name: Name,
	}
}

func (d *Entity) GetID() int {
	return 0
}
