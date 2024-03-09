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

func (s *AppTestSuite) TestShouldDepositAmountToWallet() {
	wallet, err := s.walletService.NewWallet()
	s.Nil(err)

	zeroBalance, err := s.walletService.GetBalance(wallet.ID)
	s.Nil(err)
	s.Equal("0.00", zeroBalance)

	err = s.walletService.Deposit(wallet.ID, "100.00")
	s.Nil(err)

	balance, err := s.walletService.GetBalance(wallet.ID)
	s.Nil(err)
	s.Equal("100.00", balance)
}

func (s *AppTestSuite) TestShouldGetBalance() {
	wallet, err := s.walletService.NewWallet()
	s.Nil(err)

	balance, err := s.walletService.GetBalance(wallet.ID)
	s.Nil(err)
	s.Equal("0.00", balance)
}

func (s *AppTestSuite) TestShouldGetProperBalanceAfterUpdate() {
	wallet, err := s.walletService.NewWallet()
	s.Nil(err)

	err = s.walletService.Deposit(wallet.ID, "100.00")
	s.Nil(err)

	balance, err := s.walletService.GetBalance(wallet.ID)
	s.Nil(err)
	s.Equal("100.00", balance)
}

func (s *AppTestSuite) TestShouldWithdrawAmountFromWallet() {
	wallet, err := s.walletService.NewWallet()
	s.Nil(err)

	err = s.walletService.Deposit(wallet.ID, "100.00")
	s.Nil(err)

	err = s.walletService.Withdraw(wallet.ID, "50.00")
	s.Nil(err)

	balance, err := s.walletService.GetBalance(wallet.ID)
	s.Nil(err)
	s.Equal("50.00", balance)
}

func (s *AppTestSuite) TestShouldReturnErrorWhenWithdrawMoreThanBalance() {
	wallet, err := s.walletService.NewWallet()
	s.Nil(err)

	err = s.walletService.Deposit(wallet.ID, "100.00")
	s.Nil(err)

	err = s.walletService.Withdraw(wallet.ID, "150.00")
	s.NotNil(err)
}

func TestWalletAppTestsSuite(t *testing.T) {
	suite.Run(t, new(AppTestSuite))
}
