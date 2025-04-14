package example_resources

import (
	example_data "rz-server/internal/app/example/application/example/data"

	"github.com/google/uuid"
)

type ResourceData struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateExampleMapper struct {
	data *example_data.ExampleData
}

func NewCreateExampleMapper(data *example_data.ExampleData) *CreateExampleMapper {
	return &CreateExampleMapper{data}
}

func (m *CreateExampleMapper) ToResource() ResourceData {
	return ResourceData{
		Id:   m.data.ID,
		Name: m.data.Name,
	}
}
