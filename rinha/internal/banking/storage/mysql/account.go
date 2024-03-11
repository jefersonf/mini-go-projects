package mysql

import (
	"database/sql"
	"rinha/internal/banking"
	"time"
)

const (
	accountsQueryStr         = `SELECT bal, lim FROM accounts WHERE id = ?;`
	accountsUpdateBalanceStr = `UPDATE accounts SET bal = ? WHERE id = ?;`

	retryTimeout = 3 * time.Second
)

type accountStorage struct {
	db *sql.DB
}

func (s *accountStorage) FindOne(accountID int) (banking.Account, error) {
	var a banking.Account
	row := s.db.QueryRow(accountsQueryStr, accountID)
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

retryToConnect:
	db, err := sql.Open(driverName, cfg.FormatDSN())
	if err != nil {
		time.Sleep(retryTimeout)
		goto retryToConnect
	}

	err = db.Ping()
	if err != nil {
		goto retryToConnect
	}

	return &accountStorage{db}, nil
}
