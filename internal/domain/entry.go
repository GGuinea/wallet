package domain

import "time"

type Entry struct {
	ID           string
	Type         string
	Amount       Money
	BalanceAfter Money
	CreatedAt    time.Time
}
