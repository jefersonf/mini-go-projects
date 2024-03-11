package service

import (
	"context"
	"errors"
	"fmt"
	"rinha/internal/banking"
	"sync"
	"time"
)

var (
	ErrContext           = errors.New("account service error")
	ErrInsufficientLimit = errors.New("account insufficient limit error")
)

type accountService struct {
	accounts     banking.AccountRepository
	transactions banking.TransactionRepository
	mu           sync.RWMutex
}

func NewAccountService(
	accountRepo banking.AccountRepository,
	transactionsRepository banking.TransactionRepository) *accountService {

	service := &accountService{
		accounts:     accountRepo,
		transactions: transactionsRepository,
		mu:           sync.RWMutex{},
	}

	return service
}

func (s *accountService) Credit(
	ctx context.Context,
	accountID int,
	amount int64,
	description string) (ResponseTransaction, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	account, err := s.accounts.FindOne(accountID)
	if err != nil {
		return ResponseTransaction{}, fmt.Errorf("[1]%w: %w", ErrContext, err)
	}

	newBalance, err := s.accounts.UpdateBalance(accountID, account.Balance+amount)
	if err != nil {
		return ResponseTransaction{}, fmt.Errorf("[2]%w: %w", ErrContext, err)
	}

	tx := banking.Transaction{
		AccountID:   accountID,
		Amount:      amount,
		Type:        rune(banking.Credit),
		Description: description,
	}

	_, err = s.transactions.Save(&tx)
	if err != nil {
		return ResponseTransaction{}, fmt.Errorf("[3]%w: %w", ErrContext, err)
	}

	responseTx := ResponseTransaction{
		Balance: newBalance,
		Limit:   account.Limit,
	}

	return responseTx, nil
}

func (s *accountService) Debt(
	ctx context.Context,
	accountID int,
	amount int64,
	description string) (ResponseTransaction, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	account, err := s.accounts.FindOne(accountID)
	if err != nil {
		return ResponseTransaction{}, fmt.Errorf("[4]%w: %w", ErrContext, err)
	}

	if account.Balance-amount < int64(account.Limit*(-1)) {
		return ResponseTransaction{}, fmt.Errorf("[5]%w: %w", ErrContext, ErrInsufficientLimit)
	}

	newBalance, err := s.accounts.UpdateBalance(accountID, account.Balance-amount)
	if err != nil {
		return ResponseTransaction{}, fmt.Errorf("[6]%w: %w", ErrContext, err)
	}

	tx := banking.Transaction{
		AccountID:   accountID,
		Amount:      amount,
		Type:        rune(banking.Debt),
		Description: description,
	}

	_, err = s.transactions.Save(&tx)
	if err != nil {
		return ResponseTransaction{}, fmt.Errorf("[7]%w: %w", ErrContext, err)
	}

	responseTx := ResponseTransaction{
		Balance: newBalance,
		Limit:   account.Limit,
	}

	return responseTx, nil
}

func (s *accountService) GenerateStatement(
	ctx context.Context,
	accountID int) (ResponseAccountStatement, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	account, err := s.accounts.FindOne(accountID)
	if err != nil {
		return ResponseAccountStatement{}, fmt.Errorf("%w: %w", ErrContext, err)
	}

	accountTransactions, err := s.transactions.FindByAccountID(accountID)
	if err != nil {
		return ResponseAccountStatement{}, fmt.Errorf("%w: %w", ErrContext, err)
	}

	respRecentTxs := make([]ResponseRecentTransactions, len(accountTransactions))
	for i := range respRecentTxs {
		respRecentTxs[i].Amount = accountTransactions[i].Amount
		respRecentTxs[i].Type = string(accountTransactions[i].Type)
		respRecentTxs[i].Description = accountTransactions[i].Description
		respRecentTxs[i].Timestamp = accountTransactions[i].Timestamp
	}

	accountStmt := ResponseAccountStatement{
		Balance: ResponseBalance{
			Total:     account.Balance,
			Limit:     account.Limit,
			Timestamp: time.Now().UTC(),
		},
		RecentTransactions: respRecentTxs,
	}

	return accountStmt, nil
}
