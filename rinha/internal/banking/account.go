package banking

import (
	"context"
	"net/http"
)

// Account describes a banking account definition to be shared across packages.
type Account struct {
	Balance int64 `json:"saldo"`
	Limit   int   `json:"limite"`
}

// AccountRepository provides access to the accounts storage.
type AccountRepository interface {
	FindOne(accountID int) (Account, error)
	UpdateBalance(accountID int, newBalance int64) (int64, error)
}

type AccountHTTPHandler interface {
	TransactionHandler(ctx context.Context) http.HandlerFunc
	StatementHandler(ctx context.Context) http.HandlerFunc
}
