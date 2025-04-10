package sql_store

import (
	"context"
	"database/sql"
	"embed"
	"rz-server/internal/app/example/infra/store"
	"rz-server/internal/app/example/infra/store/sql/repository"
	"rz-server/internal/common/interfaces"

	"github.com/pressly/goose/v3"
)

//go:embed schema/*.sql
var embedMigrations embed.FS

var _ store.ExampleStore = (*SqlStore)(nil)

type SqlStore struct {
	Db      *sql.DB
	Queries *repository.Queries
}

func New(db *sql.DB, util interfaces.LogUtil) *SqlStore {
	queries := repository.New(db)

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		util.Error("failed to set goose dialect", map[string]any{
			"error": err.Error(),
		})
	}

	if err := goose.Up(db, "sql/schema"); err != nil {
		util.Error("failed to run goose migrations", map[string]any{
			"error": err.Error(),
		})
	}

	return &SqlStore{
		Db:      db,
		Queries: queries,
	}
}

func (s *SqlStore) CreateExample(name string) store.ExampleStoreData {
	example, err := s.Queries.CreateExample(context.Background(), name)

	if err != nil {
		panic(err)
	}

	return store.ExampleStoreData{
		ID:   example.ID,
		Name: example.Name,
	}

}
