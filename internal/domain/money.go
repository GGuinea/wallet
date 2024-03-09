package domain

import "github.com/shopspring/decimal"

type Money interface {
	AddFromString(money string) error
	SubtractFromString(money string) error
	GetAsStringWithDefaultPrecision() string
	IsGreaterThanZero() bool
}

type DecimalMoney struct {
	amount decimal.Decimal
}

func NewDecimalMoney() *DecimalMoney {
	return &DecimalMoney{}
}

func NewDecimalMoneyFromString(money string) (*DecimalMoney, error) {
	amount, err := decimal.NewFromString(money)
	if err != nil {
		return nil, err
	}
	return &DecimalMoney{amount: amount}, nil
}

const defaultPrecision = 2

func (m *DecimalMoney) AddFromString(money string) error {
	amount, err := decimal.NewFromString(money)
	if err != nil {
		return  err
	}
	m.amount = m.amount.Add(amount)
	return nil
}

func (m *DecimalMoney) SubtractFromString(money string) error {
	amount, err := decimal.NewFromString(money)
	if err != nil {
		return err
	}
	m.amount = m.amount.Sub(amount)
	return nil
}

func (m *DecimalMoney) GetAsStringWithDefaultPrecision() string {
	return m.amount.StringFixedBank(defaultPrecision)
}

func (m *DecimalMoney) IsGreaterThanZero() bool {
	return !m.amount.IsNegative()
}
