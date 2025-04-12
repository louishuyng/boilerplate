package example_sql_store

import (
	"context"
	"database/sql"
	"embed"
	"rz-server/internal/app/example/infra/store"
	example_store_data "rz-server/internal/app/example/infra/store/sql/example/data"
	repository "rz-server/internal/app/example/infra/store/sql/example/repository"
	"rz-server/internal/common/interfaces"

	"github.com/pressly/goose/v3"
)

//go:embed schema/*.sql
var embedMigrations embed.FS

var _ store.ExampleStore = (*ExampleStore)(nil)

type ExampleStore struct {
	Queries *repository.Queries
}

func New(db *sql.DB, util *interfaces.Util) *ExampleStore {
	queries := repository.New(db)

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		util.Log.Error("failed to set goose dialect", map[string]any{
			"error": err.Error(),
		})
	}

	if err := goose.Up(db, "schema"); err != nil {
		util.Log.Error("failed to run goose migrations", map[string]any{
			"error": err.Error(),
		})
	}

	return &ExampleStore{
		Queries: queries,
	}
}

func (s *ExampleStore) CreateExample(name string) example_store_data.Data {
	example, err := s.Queries.CreateExample(context.Background(), name)

	if err != nil {
		panic(err)
	}

	return example_store_data.New(example.ID, example.Name)
}
