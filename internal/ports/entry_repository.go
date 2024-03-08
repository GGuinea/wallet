package ports

import "main/internal/domain"

type EntryRepository interface {
	Save(entry *domain.Entry) error
	GetByWalletID(walletID string) ([]*domain.Entry, error)
}
