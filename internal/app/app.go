package app

import (
	"main/internal/ports"
)

type WalletService interface {
	Deposit(walletID string, amount string) error
	Withdraw(walletID string, amount string) error
	GetBalance(walletID string) (string, error)
}

type walletService struct {
	walletRepo ports.WalletRepository
	entryRepo  ports.EntryRepository
}

func NewWalletService(walletRepo ports.WalletRepository, entryRepo ports.EntryRepository) WalletService {
	return walletService{
		walletRepo: walletRepo,
		entryRepo:  entryRepo,
	}
}

func (s walletService) Deposit(walletID string, amount string) error {
	return nil
}

func (s walletService) Withdraw(walletID string, amount string) error {
	return nil
}

func (s walletService) GetBalance(walletID string) (string, error) {
	return "", nil
}
