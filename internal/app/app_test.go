package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnNewWalletWithBalanceZero(t *testing.T) {
	walletService := NewWalletService(nil, nil)
	returnedWallet := walletService.NewWallet()
	assert.NotNil(t, returnedWallet)
	assert.Equal(t, "0.00", returnedWallet.Balance.GetAsStringWithDefaultPrecision())
}
