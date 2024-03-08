package adapters

import (
	"fmt"
	"main/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type WalletPostgresRepo struct {
	ConnPool *pgxpool.Pool
}

type WalletPostgresRepoDeps struct {
	ConnPool *pgxpool.Pool
}

func NewWalletPostgresRepo(deps *WalletPostgresRepoDeps) (*WalletPostgresRepo, error) {
	if deps == nil {
		return nil, fmt.Errorf("WalletPostgresRepoDeps is required")
	}

	return &WalletPostgresRepo{
		ConnPool: deps.ConnPool,
	}, nil
}

func (wpr *WalletPostgresRepo) SaveWallet(wallet *domain.Wallet) error {
	return nil
}

func (wpr *WalletPostgresRepo) GetWalletByID(id string) (*domain.Wallet, error) {
	return nil, nil
}

func (wpr *WalletPostgresRepo) UpdateWalletBalance(wallet *domain.Wallet, entry *domain.Entry) error {
	return nil
}

func (wpr *WalletPostgresRepo) GetEntriesByWalletID(walletID string) ([]*domain.Entry, error) {
	return nil, nil
}
