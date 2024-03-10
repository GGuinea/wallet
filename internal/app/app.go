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
	GetWalletByID(walletID string) (*domain.Wallet, error)
	GetWalletEntries(walletID string) ([]*domain.Entry, error)
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
		CreatedAt:    time.Now(),
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
		CreatedAt:    time.Now(),
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

func (s *basicWalletService) GetWalletByID(walletID string) (*domain.Wallet, error) {
	walletEntity, err := s.walletRepo.GetWalletByID(walletID)
	if err != nil {
		return nil, err
	}
	return domainWalletFromEntityWallet(walletEntity)
}

func (s *basicWalletService) GetWalletEntries(walletID string) ([]*domain.Entry, error) {
	entries, err := s.walletRepo.GetEntriesByWalletID(walletID)
	if err != nil {
		return nil, err
	}
	return domainEntriesFromEntityEntries(entries)
}

func domainEntriesFromEntityEntries(entries []*entity.EntryEntity) ([]*domain.Entry, error) {
	domainEntries := make([]*domain.Entry, 0, len(entries))
	for _, entry := range entries {
		amount, err := domain.NewDecimalMoneyFromString(entry.Amount)
		if err != nil {
			return nil, err
		}
		balanceAfter, err := domain.NewDecimalMoneyFromString(entry.BalanceAfter)
		if err != nil {
			return nil, err
		}
		domainEntries = append(domainEntries, &domain.Entry{
			ID:           entry.ID,
			WalletID:     entry.WalletID,
			Type:         entry.Type,
			Amount:       amount,
			BalanceAfter: balanceAfter,
			CreatedAt:    entry.CreatedAt,
		})
	}
	return domainEntries, nil
}

func domainWalletFromEntityWallet(walletEntity *entity.WalletEntity) (*domain.Wallet, error) {
	balance, err := domain.NewDecimalMoneyFromString(walletEntity.Balance)
	if err != nil {
		return nil, err
	}
	return &domain.Wallet{
		ID:      walletEntity.ID,
		Balance: balance,
		CreatedAt: walletEntity.CreatedAt,
		UpdatedAt: walletEntity.UpdatedAt,
	}, nil
}
