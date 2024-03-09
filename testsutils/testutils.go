package testsutils

import (
	"context"
	"fmt"
	"main/config"
	"main/pkg/postgres/migrations"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDbPool() *pgxpool.Pool {
	pgConnStr := "postgres://postgres:postgres@localhost:5432/postgres"
	pool, err := pgxpool.New(context.Background(), pgConnStr)
	if err != nil {
		panic(err)
	}
	return pool
}


func SetupDbForTests() {
	testConfig := &config.DbConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DbName:   "postgres",

	}
	err := migrations.MigrateGlobal(context.Background(), testConfig)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
