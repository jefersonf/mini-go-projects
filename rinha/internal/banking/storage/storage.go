package storage

import (
	"rinha/internal/banking"
	"rinha/internal/banking/storage/mysql"
	"rinha/internal/banking/storage/postgres"
)

type Type int

const (
	Memory Type = iota
	Postgres
	MySQL
)

func NewAccountStorage(options ...Option) banking.AccountRepository {
	var config Config
	for _, opt := range options {
		if err := opt(&config); err != nil {
			panic(err)
		}
	}

	var (
		repo banking.AccountRepository
		err  error
	)

	switch config.repositoryType {
	case Postgres:
		repo, err = postgres.NewAccountStorage()
		if err != nil {
			panic(err)
		}
	case MySQL:
		repo, err = mysql.NewAccountStorage()
		if err != nil {
			panic(err)
		}
	case Memory:
		fallthrough
	default:
		repo = nil // FIX
	}

	return repo
}

func NewTransactionStorage(options ...Option) banking.TransactionRepository {
	var config Config
	for _, opt := range options {
		if err := opt(&config); err != nil {
			panic(err)
		}
	}

	var (
		repo banking.TransactionRepository
		err  error
	)

	switch config.repositoryType {
	case Postgres:
		repo, err = postgres.NewTransactionStorage()
		if err != nil {
			panic(err)
		}
	case MySQL:
		repo, err = mysql.NewTransactionStorage()
		if err != nil {
			panic(err)
		}
	case Memory:
		fallthrough
	default:
		repo = nil // FIX
	}

	return repo
}

type Option func(*Config) error

type Config struct {
	repositoryType Type
}

func InMemory() Option {
	return func(sc *Config) error {
		sc.repositoryType = Memory
		return nil
	}
}

func WithPostgres() Option {
	return func(sc *Config) error {
		sc.repositoryType = Postgres
		return nil
	}
}

func WithMySQL() Option {
	return func(sc *Config) error {
		sc.repositoryType = MySQL
		return nil
	}
}
