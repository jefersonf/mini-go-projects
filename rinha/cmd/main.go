package main

import (
	"flag"
	"rinha/internal/api/rest"
	"rinha/internal/banking/service"
	"rinha/internal/banking/storage"
)

var defaultListenAddr string = ":8080"

func main() {

	listenAddr := flag.String("listenaddr", defaultListenAddr, "api listen address")
	flag.Parse()

	var (
		// repositories
		accountRepo      = storage.NewAccountStorage(storage.WithMySQL())
		transactionsRepo = storage.NewTransactionStorage(storage.WithMySQL())
		// services
		accountService = service.NewAccountService(accountRepo, transactionsRepo)
	)
	rest.StartNewServer(*listenAddr, accountService)
}
