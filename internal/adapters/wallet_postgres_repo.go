package adapters

import (
	"context"
	"fmt"
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

var getWalletBalanceQuery = `SELECT balance FROM wallet WHERE id = $1`

var insertEntryQuery = `INSERT INTO entry (id, wallet_id, op_type, amount, balance_after, created_at) VALUES ($1, $2, $3, $4, $5, $6)`

var getEntriesByWalletIDQuery = `SELECT * FROM entry WHERE wallet_id = $1`

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

func (wpr *WalletPostgresRepo) GetWalletBalance(id string) (string, error) {
	ctx := context.Background()
	var balance string
	err := wpr.ConnPool.QueryRow(ctx, getWalletBalanceQuery, id).Scan(&balance)

	if err != nil {
		return "", err
	}

	if balance == "" {
		return "", fmt.Errorf("wallet not found")
	}

	return balance, nil
}

func (wpr *WalletPostgresRepo) UpdateWalletBalance(wallet *entity.WalletEntity, entry *entity.EntryEntity) error {
	ctx := context.Background()
	tx, err := wpr.ConnPool.Begin(ctx)
	if err != nil {
		return err
	}

	var savedWallet entity.WalletEntity
	err = tx.QueryRow(ctx, getWalletByIDQuery, wallet.ID).Scan(&savedWallet.ID, &savedWallet.Balance, &savedWallet.CreatedAt, &savedWallet.UpdatedAt)

	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	res, err := tx.Exec(ctx, `UPDATE wallet SET balance = $1, updated_at = $2 WHERE id = $3 AND updated_at = $4 AND balance = $5`, wallet.Balance, wallet.UpdatedAt, wallet.ID, savedWallet.UpdatedAt, savedWallet.Balance)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	if res.RowsAffected() == 0 {
		tx.Rollback(ctx)
		return fmt.Errorf("wallet balance update failed")
	}

	_, err = tx.Exec(ctx, insertEntryQuery, entry.ID, entry.WalletID, entry.Type, entry.Amount, entry.BalanceAfter, entry.CreatedAt)

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

func (wpr *WalletPostgresRepo) GetEntriesByWalletID(walletID string) ([]*entity.EntryEntity, error) {
	ctx := context.Background()
	rows, err := wpr.ConnPool.Query(ctx, getEntriesByWalletIDQuery, walletID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	entries := make([]*entity.EntryEntity, 0)
	for rows.Next() {
		var entry entity.EntryEntity
		err = rows.Scan(&entry.ID, &entry.WalletID, &entry.Type, &entry.Amount, &entry.BalanceAfter, &entry.CreatedAt)
		if err != nil {
			return nil, err
		}

		entries = append(entries, &entry)
	}

	return entries, nil
}
