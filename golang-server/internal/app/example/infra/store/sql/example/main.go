package example_sql_store

import (
	"context"
	"rz-server/internal/app/example/infra/store"
	sql_store "rz-server/internal/app/example/infra/store/sql"
	example_store_data "rz-server/internal/app/example/infra/store/sql/example/data"
	repository "rz-server/internal/app/example/infra/store/sql/repository"
)

var _ store.ExampleStore = (*ExampleStore)(nil)

type ExampleStore struct {
	Queries *repository.Queries
}

func New(store *sql_store.Repository) *ExampleStore {
	return &ExampleStore{
		Queries: store.Queries,
	}
}

func (s *ExampleStore) CreateExample(name string) example_store_data.Data {
	example, err := s.Queries.CreateExample(context.Background(), name)

	if err != nil {
		panic(err)
	}

	return example_store_data.New(example.ID, example.Name)
}
