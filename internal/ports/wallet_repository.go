package ports

import (
	"main/internal/domain"
	"main/internal/entity"
)

type WalletRepository interface {
	GetWalletByID(id string) (*entity.WalletEntity, error)
	SaveWallet(wallet *entity.WalletEntity) error
	UpdateWalletBalance(wallet *entity.WalletEntity, entry *domain.Entry) error
	GetEntriesByWalletID(walletID string) ([]*domain.Entry, error)
}
