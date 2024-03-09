package ports

import (
	"main/internal/entity"
)

type WalletRepository interface {
	GetWalletByID(id string) (*entity.WalletEntity, error)
	SaveWallet(wallet *entity.WalletEntity) error
	UpdateWalletBalance(wallet *entity.WalletEntity, entry *entity.EntryEntity) error
	GetEntriesByWalletID(walletID string) ([]*entity.EntryEntity, error)
	GetWalletBalance(id string) (string, error)
}
