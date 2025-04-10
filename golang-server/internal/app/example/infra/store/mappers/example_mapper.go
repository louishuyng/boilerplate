package mapper

import (
	"rz-server/internal/app/example/domain/example"
	store "rz-server/internal/app/example/infra/store"
)

func ToDBExample(entity example.ExampleEntity) store.Example {
	return store.Example{
		Name: entity.Name,
	}
}

func FromDBExample(data store.Example) example.ExampleEntity {
	return example.ExampleEntity{
		Name: data.Name,
	}
}
