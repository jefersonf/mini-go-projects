package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"time"
)

type Medication struct {
	Interval        int    `json:"interval"`
	IntervalSize    string `json:"interval_size"`
	IntervalMod     int    `json:"interval_mod,omitempty"`
	IntervalChange  int    `json:"interval_change,omitempy"`
	Quantity        int    `json:"quantity"`
	Duration        int    `json:"duration"`
	DurationUnit    string `json:"duration_unit"`
	Type            string `json:"type"`
	FirstMedication string `json:"first_medication"`
}

type Prescription map[string]Medication

type ScheduledMedicine struct {
	Name     string
	DateTime time.Time
	Type     string
	Quantity int
}

func main() {

	data, err := os.ReadFile("prescription.json")
	if err != nil {
		panic(err)
	}

	var p Prescription
	err = json.Unmarshal(data, &p)
	if err != nil {
		fmt.Println(err)
	}

	scheduled := make([]ScheduledMedicine, 0)

	for k, v := range p {

		if v.IntervalMod == 0 {
			v.IntervalMod = math.MaxInt
		}

		duration := v.Duration
		if v.DurationUnit == "day" {
			duration *= int(time.Hour) * 24
		}

		medicationDate, err := time.Parse(time.RFC3339, v.FirstMedication)
		if err != nil {
			fmt.Println(err)
		}

		medicationStartDateOf := make(map[string]time.Time)
		endMedication := medicationDate.Add(time.Duration(duration))

		for medicationDate.Before(endMedication) {

			medication := ScheduledMedicine{
				Name:     k,
				DateTime: medicationDate,
				Type:     v.Type,
				Quantity: v.Quantity,
			}

			if _, ok := medicationStartDateOf[k]; !ok {
				medicationStartDateOf[k] = medicationDate
			} else {
				elapsedDays := medicationDate.Sub(medicationStartDateOf[k]).Hours() / 24
				if elapsedDays >= float64(v.IntervalMod) {
					v.IntervalMod += v.IntervalMod
					v.Interval += v.IntervalChange
				}
			}

			intervalUnit := v.Interval
			if v.IntervalSize == "hour" {
				intervalUnit *= int(time.Hour)
			}

			medicationDate = medicationDate.Add(time.Duration(intervalUnit))
			scheduled = append(scheduled, medication)
		}
	}

	slices.SortFunc(scheduled, func(a, b ScheduledMedicine) int {
		return a.DateTime.Compare(b.DateTime)
	})

	for i, m := range scheduled {
		fmt.Printf("%03d %s %v %s %d %s\n", i, strings.ToUpper(m.Name), m.DateTime.Format(time.DateOnly), m.DateTime.Format(time.Kitchen), m.Quantity, m.Type)
	}
}
