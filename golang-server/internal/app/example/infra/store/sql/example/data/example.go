package example_store_data

import "github.com/google/uuid"

type Data struct {
	Id   uuid.UUID
	Name string
}

func New(id uuid.UUID, name string) Data {
	return Data{
		Id:   id,
		Name: name,
	}
}
