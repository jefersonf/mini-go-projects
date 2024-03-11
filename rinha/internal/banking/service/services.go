package service

import (
	"context"
	"time"
)

type AccountService interface {
	Credit(ctx context.Context, accountID int, amount int64, description string) (ResponseTransaction, error)
	Debt(ctx context.Context, accountID int, amount int64, description string) (ResponseTransaction, error)
	GenerateStatement(ctx context.Context, accountID int) (ResponseAccountStatement, error)
}

type ResponseTransaction struct {
	Balance int64 `json:"saldo"`
	Limit   int   `json:"limite"`
}

type ResponseAccountStatement struct {
	Balance            ResponseBalance              `json:"saldo"`
	RecentTransactions []ResponseRecentTransactions `json:"ultimas_transacoes"`
}

type ResponseBalance struct {
	Total     int64     `json:"total"`
	Limit     int       `json:"limite"`
	Timestamp time.Time `json:"data_extrato"`
}

type ResponseRecentTransactions struct {
	Amount      int64     `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	Timestamp   time.Time `json:"realizada_em"`
}
