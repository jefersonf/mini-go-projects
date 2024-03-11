package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rinha/internal/banking/service"
	"time"
)

const (
	readTimeout       = 1 * time.Second
	readHeaderTimeout = 1 * time.Second
	writeTimeout      = 1 * time.Second
	idleTimeout       = 30 * time.Second
	shutdownTimeout   = 3 * time.Second
)

func StartNewServer(listenAddr string, accountService service.AccountService) {

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	server := NewServer(listenAddr, accountService)

	go func() {
		log.Printf("server running at %v", listenAddr)
		log.Println(server.ListenAndServe())
	}()

	<-stopChan

	log.Println("shutting down server")
	ctx, shutdownServer := context.WithTimeout(context.Background(), shutdownTimeout)
	defer shutdownServer()
	server.Shutdown(ctx)
	log.Println("server gracefully stopped")
}

func NewServer(listenAddr string, accountService service.AccountService) *http.Server {

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root page\n"))
	})

	ctx := context.Background()
	router.HandleFunc("POST /clientes/{id}/transacoes", accountTransactionHandler(ctx, accountService))
	router.HandleFunc("GET /clientes/{id}/extrato", accountStatementHandler(ctx, accountService))

	routerWithMiddlewares := NewResponseHeader(router, "Content-Type", "application/json")

	server := &http.Server{
		Addr:              listenAddr,
		Handler:           routerWithMiddlewares,
		ReadTimeout:       readTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
	}

	return server
}
