package adapters

import (
	"main/internal/entity"
	"main/pkg"
	"main/testsutils"
	"testing"

	"github.com/stretchr/testify/suite"
)

type WalletPostgresRepoTestSuite struct {
	suite.Suite
	walletRepo *WalletPostgresRepo
	uuidGen    pkg.UUIDGenerator
}

func (suite *WalletPostgresRepoTestSuite) SetupTest() {
	testsutils.SetupDbForTests()
	walletRepo, err := NewWalletPostgresRepo(&WalletPostgresRepoDeps{ConnPool: testsutils.GetDbPool()})
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.walletRepo = walletRepo
	suite.uuidGen = pkg.NewUUIDGenerator()
}

func (s *WalletPostgresRepoTestSuite) TestShouldSaveWallet() {
	wallet := &entity.WalletEntity{
		ID:      s.uuidGen.Generate(),
		Balance: "0.00",
	}

	err := s.walletRepo.SaveWallet(wallet)
	s.Nil(err)
}

func (s *WalletPostgresRepoTestSuite) TestShouldGetWalletByID() {
	wallet := &entity.WalletEntity{
		ID:      s.uuidGen.Generate(),
		Balance: "0.00",
	}

	err := s.walletRepo.SaveWallet(wallet)
	s.Nil(err)

	returnedWallet, err := s.walletRepo.GetWalletByID(wallet.ID)
	s.Nil(err)
	s.Equal(wallet.ID, returnedWallet.ID)
	s.Equal(wallet.Balance, returnedWallet.Balance)
}

func (s *WalletPostgresRepoTestSuite) TestShouldReturnErrorWhenWalletNotFound() {
	returnedWallet, err := s.walletRepo.GetWalletByID("non-existing-wallet-id")
	s.NotNil(err)
	s.Nil(returnedWallet)
}

func (s *WalletPostgresRepoTestSuite) TestShouldReturnErrorWhenSaveWalletWithSameID() {
	wallet := &entity.WalletEntity{
		ID:      s.uuidGen.Generate(),
		Balance: "0.00",
	}

	err := s.walletRepo.SaveWallet(wallet)
	s.Nil(err)

	err = s.walletRepo.SaveWallet(wallet)
	s.NotNil(err)
}

func (s *WalletPostgresRepoTestSuite) TestShouldUpdateWalletBalance() {
	wallet := &entity.WalletEntity{
		ID:      s.uuidGen.Generate(),
		Balance: "0.00",
	}

	err := s.walletRepo.SaveWallet(wallet)
	s.Nil(err)

	wallet.Balance = "100.00"

	entryEntity := &entity.EntryEntity{
		ID:           s.uuidGen.Generate(),
		WalletID:     wallet.ID,
		Type:         "DEPOSIT",
		Amount:       "100.00",
		BalanceAfter: "100.00",
	}
	err = s.walletRepo.UpdateWalletBalance(wallet, entryEntity)
	s.Nil(err)

	returnedWallet, err := s.walletRepo.GetWalletByID(wallet.ID)
	s.Nil(err)
	s.Equal(wallet.ID, returnedWallet.ID)
	s.Equal(wallet.Balance, returnedWallet.Balance)
}

func TestWalletPostgresRepoTestSuite(t *testing.T) {
	suite.Run(t, new(WalletPostgresRepoTestSuite))
}
