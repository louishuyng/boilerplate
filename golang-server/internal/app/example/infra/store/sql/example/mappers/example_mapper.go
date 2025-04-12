package example_store_mapper

import (
	"rz-server/internal/app/example/domain/example"
	example_store_data "rz-server/internal/app/example/infra/store/sql/example/data"
)

func ToDBExample(entity example.Entity) example_store_data.Data {
	return example_store_data.New(entity.Id, entity.Name)
}

func FromDBExample(data example_store_data.Data) example.Entity {
	return example.New().Create(data.Id, data.Name)
}
