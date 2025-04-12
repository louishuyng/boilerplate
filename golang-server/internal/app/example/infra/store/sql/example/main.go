package example_sql_store

import (
	"context"
	"rz-server/internal/app/example/infra/store"
	sql_store "rz-server/internal/app/example/infra/store/sql"
)

var _ store.ExampleStore = (*ExampleStore)(nil)

type ExampleStore struct {
	sql_store.SqlStore
}

func NewExampleStore(sqlStore sql_store.SqlStore) *ExampleStore {
	return &ExampleStore{
		SqlStore: sqlStore,
	}
}

func (s *ExampleStore) CreateExample(name string) store.ExampleStoreData {
	example, err := s.Queries.CreateExample(context.Background(), name)

	if err != nil {
		panic(err)
	}

	return store.ExampleStoreData{
		ID:   example.ID,
		Name: example.Name,
	}
}
