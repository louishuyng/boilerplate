package sql_store

import (
	"database/sql"
	"embed"
	"rz-server/internal/app/user/infra/store/sql/repository"
	"rz-server/internal/common/interfaces"

	"github.com/pressly/goose/v3"
)

//go:embed schema/*.sql
var embedMigrations embed.FS

type Repository struct {
	Queries *repository.Queries
}

func NewRepository(db *sql.DB, util *interfaces.Util) *Repository {
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

	return &Repository{
		Queries: queries,
	}
}
