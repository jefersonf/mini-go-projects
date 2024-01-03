package geoloc

import (
	"errors"
	"testing"
)

func TestFetchData(t *testing.T) {
	resp, err := FetchData()
	if errors.Is(err, ErrUnableToParseData) {
		t.Errorf("Parsing fail: %v", err)
		return
	}
	if resp == nil {
		t.Error("Empty response")
	}
}
