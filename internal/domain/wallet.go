package domain

import "time"

type Wallet struct {
	ID        string
	Balance   Money
	CreatedAt time.Time
	UpdatedAt  time.Time
}
