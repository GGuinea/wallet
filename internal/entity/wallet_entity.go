package entity

import (
	"main/internal/domain"
	"time"
)

type WalletEntity struct {
	ID        string `db:"id"`
	Balance   string `db:"balance"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func WalletEntityFromDomain(wallet *domain.Wallet) *WalletEntity {
	return &WalletEntity{
		ID:      wallet.ID,
		Balance: wallet.Balance.GetAsStringWithDefaultPrecision(),
	}
}
