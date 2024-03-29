package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnMoneyFromString(t *testing.T) {
	m, err := NewDecimalMoneyFromString("100.00")
	assert.Nil(t, err)
	assert.Equal(t, "100.00", m.GetAsStringWithDefaultPrecision())
}

func TestShouldThrowErrorWhenParsingInvalidMoney(t *testing.T) {
	_, err := NewDecimalMoneyFromString("invalid")
	assert.NotNil(t, err)
}

func TestShouldReturnPropertlyIfMoneyIsGreaterEqualThanZero(t *testing.T) {
	scenarios := []struct {
		name           string
		money          string
		expectedResult bool
	}{
		{
			name:           "Should return true for 0",
			money:          "0.00",
			expectedResult: true,
		},
		{
			name:           "Should return true for 1",
			money:          "1.00",
			expectedResult: true,
		},
		{
			name:           "Should return false for -1",
			money:          "-1.00",
			expectedResult: false,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			m, _ := NewDecimalMoneyFromString(scenario.money)
			assert.Equal(t, scenario.expectedResult, m.IsGreaterEqualThanZero(), scenario.name)
		})
	}
}

func TestShouldReturnPropertlyIfMoneyIsGreaterThanZero(t *testing.T) {
	scenarios := []struct {
		name           string
		money          string
		expectedResult bool
	}{
		{
			name:           "Should return false for 0",
			money:          "0.00",
			expectedResult: false,
		},
		{
			name:           "Should return true for 1",
			money:          "1.00",
			expectedResult: true,
		},
		{
			name:           "Should return false for -1",
			money:          "-1.00",
			expectedResult: false,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			m, _ := NewDecimalMoneyFromString(scenario.money)
			assert.Equal(t, scenario.expectedResult, m.IsGreaterThanZero(), scenario.name)
		})
	}
}

func TestShouldReturnErrorWhenParsingInvalidMoney(t *testing.T) {
	_, err := NewDecimalMoneyFromString("invalid")
	assert.NotNil(t, err)
}

func TestShouldSubtractMoneyProperly(t *testing.T) {
	scenarios := []struct {
		name             string
		startingAmount   string
		amountToSubtract string
		expectedAmount   string
	}{
		{
			name:             "Should subtract 0 from 0",
			startingAmount:   "0.00",
			amountToSubtract: "0.00",
			expectedAmount:   "0.00",
		},
		{
			name:             "Should subtract 1 from 0",
			startingAmount:   "0.00",
			amountToSubtract: "1",
			expectedAmount:   "-1.00",
		},
		{
			name:             "Should subtract small from small to be equal to 0",
			startingAmount:   "0.01",
			amountToSubtract: "0.01",
			expectedAmount:   "0.00",
		},
		{
			name:             "Should subtract big from big",
			startingAmount:   "1002.87",
			amountToSubtract: "1232.13",
			expectedAmount:   "-229.26",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			m := NewDecimalMoney()
			starinngMoney, _ := NewDecimalMoneyFromString(scenario.startingAmount)
			m.Add(starinngMoney)
			moneyToSubtract, _ := NewDecimalMoneyFromString(scenario.amountToSubtract)
			m.Subtract(moneyToSubtract)
			assert.Equal(t, scenario.expectedAmount, m.GetAsStringWithDefaultPrecision(), scenario.name)
		})
	}
}

func TestShouldAddMoneyProperly(t *testing.T) {
	scenarios := []struct {
		name           string
		startingAmount string
		amountToAdd    string
		expectedAmount string
	}{
		{
			name:           "Should add 0 to 0",
			startingAmount: "0.00",
			amountToAdd:    "0.00",
			expectedAmount: "0.00",
		},
		{
			name:           "Should add 1 to 0",
			startingAmount: "0.00",
			amountToAdd:    "1",
			expectedAmount: "1.00",
		},
		{
			name:           "Should add small to small",
			startingAmount: "0.01",
			amountToAdd:    "0.01",
			expectedAmount: "0.02",
		},
		{
			name:           "Should add big to big",
			startingAmount: "1002.87",
			amountToAdd:    "1232.13",
			expectedAmount: "2235.00",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			m := NewDecimalMoney()
			startingMoney, _ := NewDecimalMoneyFromString(scenario.startingAmount)
			m.Add(startingMoney)
			moneyToAdd, _ := NewDecimalMoneyFromString(scenario.amountToAdd)
			m.Add(moneyToAdd)
			assert.Equal(t, scenario.expectedAmount, m.GetAsStringWithDefaultPrecision(), scenario.name)
		})
	}
}
