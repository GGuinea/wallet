package app

import (
	"main/internal/adapters"
	"main/pkg"
	"main/testsutils"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AppTestSuite struct {
	suite.Suite
	walletService WalletService
}

func (suite *AppTestSuite) SetupTest() {
	testsutils.SetupDbForTests()
	walletRepo, err := adapters.NewWalletPostgresRepo(&adapters.WalletPostgresRepoDeps{ConnPool: testsutils.GetDbPool()})
	if err != nil {
		suite.T().Fatal(err)
	}

	suite.walletService = NewWalletService(walletRepo, pkg.NewUUIDGenerator())
}

func (s *AppTestSuite) TestShouldReturnNewWalletWithBalanceZero() {
	returnedWallet, err := s.walletService.NewWallet()
	s.Nil(err)
	s.NotNil(returnedWallet)
	s.Equal("0.00", returnedWallet.Balance.GetAsStringWithDefaultPrecision())
}

func TestWalletAppTestsSuite(t *testing.T) {
	suite.Run(t, new(AppTestSuite))
}
