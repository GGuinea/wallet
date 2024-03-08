package main

import (
	"context"
	"main/config"
	"main/pkg/postgres/migrations"
)

func main() {
	ctx := context.Background()

	cfg := config.NewConfig()
	err := migrations.MigrateGlobal(ctx, cfg.DbConfig)

	if err != nil {
		panic(err)
	}
}
