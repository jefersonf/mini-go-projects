package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"
)

var (
	ytHashtagEndpoint = "https://www.youtube.com/hashtag/"

	hashtag string
	topn    uint
)

type DraftDTO struct {
	Runs []struct {
		Text string `json:"text"`
	} `json:"runs"`
	Accessibility struct {
		AccessibilityData struct {
			Label string `json:"label"`
		} `json:"accessibilityData"`
	} `json:"accessibility"`
}

type VideoInfo struct {
	ChannelName string
	VideoTitle  string
	Views       int
	Year        int
}

func main() {

	flag.StringVar(&hashtag, "ht", "", "target hashtag")
	flag.UintVar(&topn, "t", 10, "top N")
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

	type cut struct {
		L int
		R int
	}

	cuts := make([]cut, 0)
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
					cuts = append(cuts, cut{L: i + 8, R: i + j + 1})
				}
				break
			}
			j++
		}
		i = i + j + 1
	}

	rank := make([]VideoInfo, 0)
	for i := range cuts {
		var d DraftDTO
		err = json.Unmarshal(body[cuts[i].L:cuts[i].R], &d)
		if err != nil {
			continue
		}
		title := d.Runs[0].Text
		label := d.Accessibility.AccessibilityData.Label
		r1 := strings.Split(label[len(title)+1:], " ")

		viewsIdx, views := len(r1)-1, 0
		for i, s := range r1 {
			if x, err := strconv.Atoi(strings.ReplaceAll(s, ".", "")); err == nil {
				if i < viewsIdx {
					viewsIdx = i
					views = x
				}
			}
		}

		channelName := strings.Join(r1[:viewsIdx], " ")
		yearPreffixPattern := "visualizações há"
		yearIdx := strings.Index(label, yearPreffixPattern) + len(yearPreffixPattern)

		yearStr := strings.Split(strings.TrimSpace(label[yearIdx:]), " ")[0]
		year, _ := strconv.Atoi(yearStr)

		rank = append(rank, VideoInfo{
			ChannelName: channelName,
			VideoTitle:  title,
			Views:       views,
			Year:        time.Now().Year() - year,
		})
	}

	slices.SortFunc(rank, func(a, b VideoInfo) int {
		if a.Views == b.Views {
			return a.Year - b.Year
		}
		return b.Views - a.Views
	})

	fmt.Printf("Top %v most popular videos on Youtube with #%s\n\n", topn, hashtag)

	for i, v := range rank {
		if i >= int(topn) {
			break
		}
		fmt.Printf("%2d\tTitle:   %s\n\tViews:   %v\n\tChannel: %s\n\tYear:    %v\t\n\n", i+1, v.VideoTitle, v.Views, v.ChannelName, v.Year)
	}
}
