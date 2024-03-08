package domain

import "github.com/shopspring/decimal"

type Money interface {
	AddFromString(money string) (Money, error)
	SubtractFromString(money string) (Money, error)
	GetAsStringWithDefaultPrecision() string
}

type DecimalMoney struct {
	amount decimal.Decimal
}

func NewDecimalMoney() DecimalMoney {
	return DecimalMoney{}
}

const defaultPrecision = 2

func (m *DecimalMoney) AddFromString(money string) (Money, error) {
	amount, err := decimal.NewFromString(money)
	if err != nil {
		return nil, err
	}
	m.amount = m.amount.Add(amount)
	return m, nil
}

func (m *DecimalMoney) SubtractFromString(money string) (Money, error) {
	amount, err := decimal.NewFromString(money)
	if err != nil {
		return nil, err
	}
	m.amount = m.amount.Sub(amount)
	return m, nil
}

func (m *DecimalMoney) GetAsStringWithDefaultPrecision() string {
	return m.amount.StringFixedBank(defaultPrecision)
}
