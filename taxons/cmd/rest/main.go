package rest

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/jefersonf/taxons/api"
)

func main() {
	listenAddr := flag.String("listenaddr", "localhost:8080", "Taxon API listen address")
	appLang := flag.String("lang", "en", "taxon app language")
	flag.Parse()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	APIServer := api.ListenAndServe(*listenAddr, *appLang)

	<-stopChan
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	APIServer.Shutdown(ctx)
	defer cancel()
	log.Println("API server stopped!")
}
