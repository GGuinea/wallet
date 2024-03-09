package app

import (
	"main/internal/adapters"
	"main/pkg"
	"main/testsutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnNewWalletWithBalanceZero(t *testing.T) {
	walletRepo, _ := adapters.NewWalletPostgresRepo(&adapters.WalletPostgresRepoDeps{ConnPool: testsutils.GetDbPool()})
	walletService := NewWalletService(walletRepo, pkg.NewUUIDGenerator())
	returnedWallet, err := walletService.NewWallet()
	assert.Nil(t, err)
	assert.NotNil(t, returnedWallet)
	assert.Equal(t, "0.00", returnedWallet.Balance.GetAsStringWithDefaultPrecision())
}
