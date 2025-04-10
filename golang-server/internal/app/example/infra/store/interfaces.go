package store

import "github.com/google/uuid"

type Example struct {
	ID   uuid.UUID
	Name string
}

type ExampleStore interface {
	CreateExample(name string) Example
}
