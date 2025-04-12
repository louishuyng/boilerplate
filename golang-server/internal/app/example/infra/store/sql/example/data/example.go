package example_store_data

import "github.com/google/uuid"

type ExampleStoreData struct {
	Id   uuid.UUID
	Name string
}

func New(id uuid.UUID, name string) ExampleStoreData {
	return ExampleStoreData{
		Id:   id,
		Name: name,
	}
}
