package domain

import "time"

type Entry struct {
	ID           string
	WalletID     string
	Type         string
	Amount       Money
	BalanceAfter Money
	CreatedAt    time.Time
}
