package inmet

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const (
	ActiveAlertsEndpoint = "https://apiprevmet3.inmet.gov.br/avisos/ativos"
	// alertEndpoint = "https://apiprevmet3.inmet.gov.br/aviso/getByID/%v"
)

func FetchData() (*ActiveAlerts, error) {

	var alerts ActiveAlerts

	log.Println("Fetching data from INMET..")
	resp, err := http.Get(ActiveAlertsEndpoint)
	if err != nil {
		return &alerts, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &alerts, err
	}

	err = json.Unmarshal(data, &alerts)
	if err != nil {
		return &alerts, err
	}

	return &alerts, nil
}
