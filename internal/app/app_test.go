package app

import (
	"main/internal/adapters"
	"main/internal/domain"
	"main/pkg"
	"main/testsutils"
	"sync"
	"testing"
	"time"

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

func (s *AppTestSuite) TestShouldReturnErrorWhenWithdrawNegativeAmount() {
	wallet, err := s.walletService.NewWallet()
	s.Nil(err)

	err = s.walletService.Deposit(wallet.ID, "100.00")
	s.Nil(err)

	err = s.walletService.Withdraw(wallet.ID, "-50.00")
	s.NotNil(err)
}

func (s *AppTestSuite) TestShouldReturnErrorWhenDepositNegativeAmount() {
	wallet, err := s.walletService.NewWallet()
	s.Nil(err)

	err = s.walletService.Deposit(wallet.ID, "-100.00")
	s.NotNil(err)
}

func (s *AppTestSuite) TestShouldReturnErrorWhenDepositZeroAmount() {
	wallet, err := s.walletService.NewWallet()
	s.Nil(err)

	err = s.walletService.Deposit(wallet.ID, "0.00")
	s.NotNil(err)
}

func (s *AppTestSuite) TestShouldReturnErrorWhenWithdrawZeroAmount() {
	wallet, err := s.walletService.NewWallet()
	s.Nil(err)

	err = s.walletService.Withdraw(wallet.ID, "0.00")
	s.NotNil(err)
}

func (s *AppTestSuite) TestShouldReturnErrorWhenWithdrawInvalidAmount() {
	wallet, err := s.walletService.NewWallet()
	s.Nil(err)

	err = s.walletService.Withdraw(wallet.ID, "invalid")
	s.NotNil(err)
}

func (s *AppTestSuite) TestShouldReturnErrorWhenDepositInvalidAmount() {
	wallet, err := s.walletService.NewWallet()
	s.Nil(err)

	err = s.walletService.Deposit(wallet.ID, "invalid")
	s.NotNil(err)
}

func (s *AppTestSuite) TestShouldReturnErrorWhenGetBalanceForInvalidWallet() {
	_, err := s.walletService.GetBalance("invalid")
	s.NotNil(err)
}

func (s *AppTestSuite) TestShouldReturnErrorWhenDepositToInvalidWallet() {
	err := s.walletService.Deposit("invalid", "100.00")
	s.NotNil(err)
}

func (s *AppTestSuite) TestShouldReturnErrorWhenWithdrawFromInvalidWallet() {
	err := s.walletService.Withdraw("invalid", "100.00")
	s.NotNil(err)
}

func (s *AppTestSuite) TestShouldHandleConcurrentWithdraws() {
	wallet, err := s.walletService.NewWallet()
	s.Nil(err)

	err = s.walletService.Deposit(wallet.ID, "12.00")
	s.Nil(err)

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		go func() {
			defer wg.Done()
			_ = s.walletService.Withdraw(wallet.ID, "2.60")
		}()
	}
	wg.Wait()

	balance, err := s.walletService.GetBalance(wallet.ID)
	s.Nil(err)
	money, err := domain.NewDecimalMoneyFromString(balance)
	s.Nil(err)
	s.True(money.IsGreaterEqualThanZero())
}

func TestWalletAppTestsSuite(t *testing.T) {
	suite.Run(t, new(AppTestSuite))
}
