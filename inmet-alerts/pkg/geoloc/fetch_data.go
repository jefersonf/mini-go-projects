package geoloc

import (
	"encoding/json"
	"errors"
	"net/http"
)

const APIEndpoint = "https://ipapi.co/json/"

var (
	ErrUnableToFetchData = errors.New("unable to get data from ipapi.co")
	ErrUnableToParseData = errors.New("unable to parse ipapi.co data")
)

func FetchData() (*GeoLocationInfo, error) {

	var geoLocInfo GeoLocationInfo

	resp, err := http.Get(APIEndpoint)
	if err != nil {
		return &geoLocInfo, ErrUnableToFetchData
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&geoLocInfo)
	if err != nil {
		return &geoLocInfo, ErrUnableToParseData
	}

	return &geoLocInfo, nil
}
