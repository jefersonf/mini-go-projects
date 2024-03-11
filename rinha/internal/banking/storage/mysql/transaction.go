package mysql

import (
	"database/sql"
	"rinha/internal/banking"
)

const (
	txInsertStr = `INSERT INTO transactions (acc_id, amt, typ, des) VALUES (?, ?, ?, ?);`
	txQueryStr  = `SELECT amt, typ, des, ts FROM transactions WHERE acc_id = ? ORDER BY ts DESC LIMIT 10;`
)

type transactionStorage struct {
	db *sql.DB
}

func (s *transactionStorage) Save(t *banking.Transaction) (int64, error) {
	result, err := s.db.Exec(txInsertStr, &t.AccountID, &t.Amount, string(t.Type), &t.Description)
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
		var typ string
		err := rows.Scan(&t.Amount, &typ, &t.Description, &t.Timestamp)
		if err != nil {
			return transactions, err
		}
		t.Type = rune(typ[0])
		// t.Timestamp = time
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

	return &transactionStorage{db}, nil
}
