package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
)

var (
	ytHashtagEndpoint = "https://www.youtube.com/hashtag/"

	hashtag string
)

type Data struct {
	Runs []struct {
		Text string `json:"text"`
	} `json:"runs"`
	Accessibility struct {
		AccessibilityData struct {
			Label string `json:"label"`
		} `json:"accessibilityData"`
	} `json:"accessibility"`
}

func main() {

	flag.StringVar(&hashtag, "ht", "", "target hashtag")
	flag.Parse()

	if len(hashtag) < 2 {
		fmt.Println(errors.New("undefined hashtag"))
	}

	url := fmt.Sprintf("%s/%s", ytHashtagEndpoint, hashtag)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	type Cut struct {
		L int
		R int
	}

	cuts := make([]Cut, 0)
	pattern := []byte(`"title":{"runs":[{"text":`)

	i := 0
	for i < len(body)-100 {
		p := i
		i = bytes.Index(body[i:], pattern)
		i += p
		if i < 0 {
			break
		}
		count := 4 // magic trick
		j := 0
		for _, r := range body[i:] {
			if rune(r) == '}' {
				count--
			}
			if count == 0 {
				if i+8 < i+j+1 {
					cuts = append(cuts, Cut{L: i + 8, R: i + j + 1})
				}
				break
			}
			j++
		}
		i = i + j + 1
	}

	for i := range cuts {
		var d Data
		err = json.Unmarshal(body[cuts[i].L:cuts[i].R], &d)
		if err != nil {
			continue
		}
		fmt.Println(d.Accessibility.AccessibilityData.Label)
	}

}
