package ports

import "main/internal/domain"

type WalletRepository interface {
	GetByID(id string) (*domain.Wallet, error)
	UpdateBalance(wallet *domain.Wallet) error
}
