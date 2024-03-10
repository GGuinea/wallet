package conn

import (
	"context"
	"fmt"
	"main/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDbPool(ctx context.Context, config *config.DbConfig) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, buildConnectionString(config))
	if err != nil {
		return nil, err

	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func buildConnectionString(dbConfig *config.DbConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
}
