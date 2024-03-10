package main

import (
	"context"
	"main/api"
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

	httpServer, err := api.NewGinHttpServer(&api.GinServerDeps{Config: cfg})
	if err != nil {
		panic(err)
	}

	err = httpServer.Start()
	if err != nil {
		panic(err)
	}
}
