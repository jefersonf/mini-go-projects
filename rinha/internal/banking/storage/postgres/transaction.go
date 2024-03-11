package postgres

import (
	"database/sql"
	"rinha/internal/banking"
)

const (
	txInsertStr = `INSERT INTO transactions (acc_id, amt, typ, des) VALUES ($1, $2, $3);`
	txQueryStr  = `SELECT amt, typ, des, ts FROM transactions WHERE acc_id = $1;`
)

type transactionStorage struct {
	db *sql.DB
}

func (s *transactionStorage) Save(t *banking.Transaction) (int64, error) {
	result, err := s.db.Exec(txInsertStr, &t.AccountID, &t.Amount, &t.Type, &t.Description)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *transactionStorage) FindByAccountID(accountID int) ([]banking.Transaction, error) {
	transactions := make([]banking.Transaction, 0)
	rows, err := s.db.Query(txQueryStr, accountID)
	if err != nil {
		return transactions, err
	}
	var t banking.Transaction
	for rows.Next() {
		err := rows.Scan(&t.Amount, &t.Type, &t.Description, &t.Timestamp)
		if err != nil {
			return transactions, err
		}
		transactions = append(transactions, t)
	}

	if rows.Err() != nil {
		return []banking.Transaction{}, err
	}

	return transactions, nil
}

func NewTransactionStorage(configs ...ConfigOption) (*transactionStorage, error) {
	cfg := NewConfig()
	for _, setConfig := range configs {
		err := setConfig(&cfg)
		if err != nil {
			return nil, err
		}
	}

	db, err := sql.Open(driverName, cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return &transactionStorage{db}, nil
}
