package conn

import (
	"context"
	"fmt"
	"main/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDbPool(ctx context.Context, config *config.DbConfig) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, buildConnectionString(config))
}

func buildConnectionString(dbConfig *config.DbConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
}
