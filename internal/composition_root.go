package internal

import (
	"context"
	"main/config"
	"main/internal/adapters"
	"main/internal/ports"
	"main/pkg/postgres/conn"
)

type CompositionRoot struct {
	WalletRepository ports.WalletRepository
}

func NewCompositionRoot(cfg *config.Config) *CompositionRoot {
	dbPool, err := conn.GetDbPool(context.Background(), cfg.DbConfig)
	if err != nil {
		panic(err)
	}

	postgresRepoDeps := adapters.WalletPostgresRepoDeps{
		ConnPool: dbPool,
	}
	posgresRepo, err := adapters.NewWalletPostgresRepo(&postgresRepoDeps)
	if err != nil {
		panic(err)
	}

	return &CompositionRoot{
		WalletRepository: posgresRepo,
	}
}
