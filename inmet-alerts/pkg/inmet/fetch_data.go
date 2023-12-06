package inmet

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

const ActiveAlertsEndpoint = "https://apiprevmet3.inmet.gov.br/avisos/ativos"

var (
	ErrUnableToFetchDataFromINMET = errors.New("unable to get data from INMET")
	ErrUnableToReadINMETData      = errors.New("unable to read INMET data")
	ErrUnableToParseINMETData     = errors.New("unable to parse INMET data")
)

func FetchData() (*ActiveAlerts, error) {

	var alerts ActiveAlerts

	log.Println("Fetching data from INMET..")
	resp, err := http.Get(ActiveAlertsEndpoint)
	if err != nil {
		return &alerts, ErrUnableToFetchDataFromINMET
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &alerts, ErrUnableToReadINMETData
	}

	err = json.Unmarshal(data, &alerts)
	if err != nil {
		return &alerts, ErrUnableToParseINMETData
	}

	return &alerts, nil
}
