package ports

import "main/internal/domain"

type WalletRepository interface {
	GetWalletByID(id string) (*domain.Wallet, error)
	SaveWallet(wallet *domain.Wallet) error
	UpdateWalletBalance(wallet *domain.Wallet, entry *domain.Entry) error
	GetEntriesByWalletID(walletID string) ([]*domain.Entry, error)
}
