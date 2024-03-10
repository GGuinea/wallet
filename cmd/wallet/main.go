package main

import (
	"context"
	"log"
	"main/api"
	"main/config"
	"main/pkg/postgres/migrations"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	go func() {
		if err := httpServer.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	contextWithTimeout, cancel := context.WithTimeout(ctx, time.Duration(cfg.ServerConfig.ShoutdownTimeout)*time.Second)
	defer cancel()

	if err := httpServer.GracefulStop(contextWithTimeout); err != nil {
		log.Println("Server shoutdown error", err)
	}

	select {
	case <-contextWithTimeout.Done():
		log.Println("Server shoutdown")
	}
}
