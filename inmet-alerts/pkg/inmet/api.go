package inmet

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	ActiveAlertsEndpoint = "https://apiprevmet3.inmet.gov.br/avisos/ativos"

	alertEndpoint = "https://apiprevmet3.inmet.gov.br/aviso/getByID/%v"
)

func FetchData() {

	resp, err := http.Get(fmt.Sprintf(alertEndpoint, 45678))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var alert BaseAlert

	err = json.Unmarshal(data, &alert)
	if err != nil {
		fmt.Println("error while unmarchaling: ", err)
	}

	fmt.Printf("%#+v\n", alert)
}
