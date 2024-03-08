package app

import (
	"main/internal/domain"
	"main/internal/ports"
	"main/pkg"
)

type basicWalletService struct {
	walletRepo ports.WalletRepository
	idGen      pkg.UUIDGenerator
}

func NewWalletService(walletRepo ports.WalletRepository, idGen pkg.UUIDGenerator) basicWalletService {
	return basicWalletService{
		walletRepo: walletRepo,
		idGen:      idGen,
	}
}

func (s *basicWalletService) NewWallet() *domain.Wallet {
	return &domain.Wallet{
		ID:      s.idGen.Generate(),
		Balance: domain.NewDecimalMoney(),
	}
}

func (s *basicWalletService) Deposit(walletID string, amount string) error {
	return nil
}

func (s *basicWalletService) Withdraw(walletID string, amount string) error {
	return nil
}

func (s *basicWalletService) GetBalance(walletID string) (string, error) {
	return "", nil
}
