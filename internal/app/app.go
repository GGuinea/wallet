package app

import (
	"fmt"
	"main/internal/domain"
	"main/internal/domain/errors"
	"main/internal/entity"
	"main/internal/ports"
	"main/pkg"
	"time"
)

type WalletService interface {
	NewWallet() (*domain.Wallet, error)
	Deposit(walletID string, amount string) error
	Withdraw(walletID string, amount string) error
	GetBalance(walletID string) (string, error)
}

type basicWalletService struct {
	walletRepo ports.WalletRepository
	idGen      pkg.UUIDGenerator
}

func NewWalletService(walletRepo ports.WalletRepository, idGen pkg.UUIDGenerator) WalletService {
	return &basicWalletService{
		walletRepo: walletRepo,
		idGen:      idGen,
	}
}

const (
	depositType  = "DEPOSIT"
	withdrawType = "WITHDRAW"
)

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
	moneyToDeposit, err := domain.NewDecimalMoneyFromString(amount)
	if err != nil {
		return err
	}

	if !moneyToDeposit.IsGreaterThanZero() {
		return errors.NewInvalidAmountError(fmt.Sprintf("invalid amount: %s", amount))
	}

	walletBalance, err := s.walletRepo.GetWalletByID(walletID)
	if err != nil {
		return err
	}

	walletMoney, err := domain.NewDecimalMoneyFromString(walletBalance.Balance)

	if err != nil {
		return err
	}

	err = walletMoney.Add(moneyToDeposit)
	if err != nil {
		return err
	}

	entryEnity := &entity.EntryEntity{
		ID:           s.idGen.Generate(),
		WalletID:     walletID,
		Type:         depositType,
		Amount:       amount,
		BalanceAfter: walletMoney.GetAsStringWithDefaultPrecision(),
	}

	newWalletEntity := &entity.WalletEntity{
		ID:        walletID,
		Balance:   walletMoney.GetAsStringWithDefaultPrecision(),
		UpdatedAt: time.Now(),
	}

	err = s.walletRepo.UpdateWalletBalance(newWalletEntity, entryEnity)

	if err != nil {
		return err
	}

	return nil
}

func (s *basicWalletService) Withdraw(walletID string, amount string) error {
	moneyToWithdraw, err := domain.NewDecimalMoneyFromString(amount)
	if err != nil {
		return err
	}

	if !moneyToWithdraw.IsGreaterThanZero() {
		return errors.NewInvalidAmountError(fmt.Sprintf("invalid amount: %s", amount))
	}

	walletBalance, err := s.walletRepo.GetWalletByID(walletID)
	if err != nil {
		return err
	}

	walletMoney, err := domain.NewDecimalMoneyFromString(walletBalance.Balance)
	if err != nil {
		return err
	}

	err = walletMoney.Subtract(moneyToWithdraw)
	if err != nil {
		return err
	}

	if !walletMoney.IsGreaterEqualThanZero() {
		return errors.NewInsufficientFundsError(fmt.Sprintf("insufficient funds: %s", walletMoney.GetAsStringWithDefaultPrecision()))
	}

	entryEnity := &entity.EntryEntity{
		ID:           s.idGen.Generate(),
		WalletID:     walletID,
		Type:         withdrawType,
		Amount:       amount,
		BalanceAfter: walletMoney.GetAsStringWithDefaultPrecision(),
	}

	newWalletEntity := &entity.WalletEntity{
		ID:        walletID,
		Balance:   walletMoney.GetAsStringWithDefaultPrecision(),
		UpdatedAt: time.Now(),
	}

	err = s.walletRepo.UpdateWalletBalance(newWalletEntity, entryEnity)
	if err != nil {
		return err
	}

	return nil
}

func (s *basicWalletService) GetBalance(walletID string) (string, error) {
	balance, err := s.walletRepo.GetWalletBalance(walletID)
	if err != nil {
		return "", err
	}
	return balance, nil
}

func domainWalletFromEntityWallet(walletEntity *entity.WalletEntity) (*domain.Wallet, error) {
	balance, err := domain.NewDecimalMoneyFromString(walletEntity.Balance)
	if err != nil {
		return nil, err
	}
	return &domain.Wallet{
		ID:      walletEntity.ID,
		Balance: balance,
	}, nil
}
