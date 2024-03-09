package entity

import (
	"main/internal/domain"
	"time"
)

type EntryEntity struct {
	ID           string    `db:"id"`
	WalletID     string    `db:"wallet_id"`
	Type         string    `db:"op_type"`
	Amount       string    `db:"amount"`
	BalanceAfter string    `db:"balance_after"`
	CreatedAt    time.Time `db:"created_at"`
}

func EntryEntityFromDomain(entry *domain.Entry) *EntryEntity {
	return &EntryEntity{
		ID:           entry.ID,
		WalletID:     entry.WalletID,
		Type:         entry.Type,
		Amount:       entry.Amount.GetAsStringWithDefaultPrecision(),
		BalanceAfter: entry.BalanceAfter.GetAsStringWithDefaultPrecision(),
		CreatedAt:    entry.CreatedAt,
	}
}
