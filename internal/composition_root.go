package internal

import (
	"context"
	"fmt"
	"main/config"
	"main/internal/adapters"
	"main/internal/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CompositionRoot struct {
	WalletRepository ports.WalletRepository
}

func NewCompositionRoot(cfg *config.Config) *CompositionRoot {
	dbPool, err := getDbPool(context.Background(), cfg.DbConfig)
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

func getDbPool(ctx context.Context, config *config.DbConfig) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, buildConnectionString(config))
}

func buildConnectionString(dbConfig *config.DbConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
}
