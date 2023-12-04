package main

import (
	"log"

	"github.com/jefersonf/mini-go-projects/inmet-alerts/pkg/inmet"
)

func main() {
	activeAlerts, err := inmet.FetchData()
	if err != nil {
		panic(err)
	}

	today, future := activeAlerts.GetAll()
	log.Printf("Active Weather Alerts: %d, today (%d) future (%d)\n", activeAlerts.Count(), today.Count(), future.Count())
}
