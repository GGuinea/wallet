package migrations

import (
	"context"
	"embed"
	"main/config"
	"main/pkg/postgres/conn"

	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed global/*.sql
var globalMigrations embed.FS

func MigrateGlobal(ctx context.Context, cfg *config.DbConfig) error {
	pool, err := conn.GetDbPool(ctx, cfg)
	if err != nil {
		return err
	}
	goose.SetDialect("postgres")
	goose.SetBaseFS(globalMigrations)

	db := stdlib.OpenDBFromPool(pool)

	defer db.Close()

	if err := goose.Up(db, "global"); err != nil {
		return err
	}

	return nil
}
