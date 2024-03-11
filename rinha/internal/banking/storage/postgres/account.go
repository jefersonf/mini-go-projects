package postgres

import (
	"database/sql"
	"rinha/internal/banking"

	_ "github.com/lib/pq"
)

const (
	accountQueryStr          = `SELECT bal, lim FROM accounts WHERE id = $1;`
	accountsUpdateBalanceStr = `UPDATE accounts SET bal = $1 WHERE id = $2`
)

type accountStorage struct {
	db *sql.DB
}

func (s *accountStorage) FindOne(accountID int) (banking.Account, error) {
	var a banking.Account
	row := s.db.QueryRow(accountQueryStr, accountID)
	err := row.Scan(&a.Balance, &a.Limit)
	if err != nil {
		return a, err
	}
	return a, nil
}

func (s *accountStorage) UpdateBalance(accountID int, newBalance int64) (int64, error) {
	_, err := s.db.Exec(accountsUpdateBalanceStr, newBalance, accountID)
	if err != nil {
		return 0, err
	}
	return newBalance, nil
}

func NewAccountStorage(configs ...ConfigOption) (*accountStorage, error) {
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

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &accountStorage{db}, nil
}
