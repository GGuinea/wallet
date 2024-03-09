package app

import (
	"main/internal/domain"
	"main/internal/entity"
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

func (s *basicWalletService) NewWallet() (*domain.Wallet, error) {
	newWallet := domain.Wallet{
		ID:      s.idGen.Generate(),
		Balance: domain.NewDecimalMoney(),
	}

	err := s.walletRepo.SaveWallet(entity.WalletEntityFromDomain(&newWallet))
	if err != nil {
		return nil, err
	}

	return &newWallet, nil
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
