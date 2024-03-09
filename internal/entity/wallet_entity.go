package entity

import "main/internal/domain"

type WalletEntity struct {
	ID        string `db:"id" validate:"required"`
	Balance   string `db:"balance" validate:"required"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func WalletEntityFromDomain(wallet *domain.Wallet) *WalletEntity {
	return &WalletEntity{
		ID:      wallet.ID,
		Balance: wallet.Balance.GetAsStringWithDefaultPrecision(),
	}
}
