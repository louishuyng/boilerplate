package example_store_mapper

import (
	"rz-server/internal/app/example/domain/example"
	example_store_data "rz-server/internal/app/example/infra/store/sql/example/data"
)

func ToDBExample(entity example.ExampleEntity) example_store_data.ExampleStoreData {
	return example_store_data.New(entity.Id, entity.Name)
}

func FromDBExample(data example_store_data.ExampleStoreData) example.ExampleEntity {
	return example.NewEntity().Create(data.Id, data.Name)
}
