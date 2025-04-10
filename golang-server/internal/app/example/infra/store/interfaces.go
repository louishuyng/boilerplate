package store

import "github.com/google/uuid"

type ExampleStoreData struct {
	ID   uuid.UUID
	Name string
}

type ExampleStore interface {
	CreateExample(name string) ExampleStoreData
}
