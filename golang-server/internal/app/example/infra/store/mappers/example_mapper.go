package mapper

import (
	"rz-server/internal/app/example/domain/example"
	store "rz-server/internal/app/example/infra/store"
)

func ToDBExample(entity example.ExampleEntity) store.ExampleStoreData {
	return store.ExampleStoreData{
		Name: entity.Name,
	}
}

func FromDBExample(data store.ExampleStoreData) example.ExampleEntity {
	return example.ExampleEntity{
		Name: data.Name,
	}
}
