package adapters

import (
	"context"
	"fmt"
	"main/internal/domain"
	"main/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type WalletPostgresRepo struct {
	ConnPool *pgxpool.Pool
}

type WalletPostgresRepoDeps struct {
	ConnPool *pgxpool.Pool
}

var insertWalletQuery = `INSERT INTO wallet (id, balance) VALUES ($1, $2)`

var getWalletByIDQuery = `SELECT * FROM wallet WHERE id = $1`

func NewWalletPostgresRepo(deps *WalletPostgresRepoDeps) (*WalletPostgresRepo, error) {
	if deps == nil {
		return nil, fmt.Errorf("WalletPostgresRepoDeps is required")
	}

	return &WalletPostgresRepo{
		ConnPool: deps.ConnPool,
	}, nil
}

func (wpr *WalletPostgresRepo) SaveWallet(wallet *entity.WalletEntity) error {
	ctx := context.Background()
	tx, err := wpr.ConnPool.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, insertWalletQuery, wallet.ID, wallet.Balance)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (wpr *WalletPostgresRepo) GetWalletByID(id string) (*entity.WalletEntity, error) {
	ctx := context.Background()
	var wallet entity.WalletEntity
	err := wpr.ConnPool.QueryRow(ctx, getWalletByIDQuery, id).Scan(&wallet.ID, &wallet.Balance, &wallet.CreatedAt, &wallet.UpdatedAt)

	if err != nil {
		return nil, err
	}

	if wallet.ID == "" {
		return nil, fmt.Errorf("wallet not found")
	}

	return &wallet, nil
}

func (wpr *WalletPostgresRepo) UpdateWalletBalance(wallet *entity.WalletEntity, entry *domain.Entry) error {
	ctx := context.Background()
	tx, err := wpr.ConnPool.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `UPDATE wallet SET balance = $1 WHERE id = $2`, wallet.Balance, wallet.ID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (wpr *WalletPostgresRepo) GetEntriesByWalletID(walletID string) ([]*domain.Entry, error) {
	return nil, nil
}
