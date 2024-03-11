package banking

import "time"

type TxType rune

const (
	Credit TxType = 'c'
	Debt   TxType = 'd'
)

type RequestTransaction struct {
	Amount      int64  `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

// Transaction describes a transaction to be shared across all packages.
type Transaction struct {
	AccountID   int       `json:"client_id"`
	Amount      int64     `json:"valor"`
	Type        rune      `json:"tipo"`
	Description string    `json:"descricao"`
	Timestamp   time.Time `json:"realizado_em"`
}

// TransactionRepository provides access to the transactions storage.
type TransactionRepository interface {
	Save(*Transaction) (int64, error)
	FindByAccountID(accountID int) ([]Transaction, error)
}
