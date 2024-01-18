package api

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func ListenAndServe(listenAddress, appLang string) *http.Server {

	routes := http.NewServeMux()
	routes.HandleFunc("api/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("api/v1"))
	})

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	server := &http.Server{
		Addr:    listenAddress,
		Handler: routes,
	}

	go func() {
		panic(server.ListenAndServe())
	}()

	<-stopChan
	fmt.Println("API gracefully stopped")

	return server
}
