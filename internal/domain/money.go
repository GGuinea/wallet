package domain

import "github.com/shopspring/decimal"

type Money interface {
	Add(money Money) error
	Subtract(money Money) error
	GetAsStringWithDefaultPrecision() string
	IsGreaterEqualThanZero() bool
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

func (m *DecimalMoney) Subtract(money Money) error {
	decimalMoney, ok := money.(*DecimalMoney)
	if !ok {
		return nil
	}

	m.amount = m.amount.Sub(decimalMoney.amount)
	return nil
}

func (m *DecimalMoney) GetAsStringWithDefaultPrecision() string {
	return m.amount.StringFixedBank(defaultPrecision)
}

func (m *DecimalMoney) IsGreaterEqualThanZero() bool {
	return !m.amount.IsNegative()
}

func (m *DecimalMoney) IsGreaterThanZero() bool {
	return m.amount.IsPositive()
}

func (m *DecimalMoney) Add(money Money) error {
	decimalMoney, ok := money.(*DecimalMoney)
	if !ok {
		return nil
	}

	m.amount = m.amount.Add(decimalMoney.amount)
	return nil
}
