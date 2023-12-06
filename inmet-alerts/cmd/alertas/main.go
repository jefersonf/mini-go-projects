package main

import (
	"log"

	"github.com/jefersonf/mini-go-projects/inmet-alerts/pkg/geoloc"
	"github.com/jefersonf/mini-go-projects/inmet-alerts/pkg/inmet"
)

func main() {

	activeAlerts, err := inmet.FetchData()
	checkErr(err)

	todayAlerts, futureAlerts := activeAlerts.GetAll()
	log.Printf("Active weather alerts: %d, today (%d) future (%d)\n", activeAlerts.Count(), todayAlerts.Count(), futureAlerts.Count())

	geoLocInfo, err := geoloc.FetchData()
	checkErr(err)

	log.Printf("Your current location: %s, %s - %s\n", geoLocInfo.City, geoLocInfo.Region, geoLocInfo.RegionCode)

	todayAlertsByRegion := todayAlerts.ByRegion(geoLocInfo.Region)
	futureAlertsByRegion := futureAlerts.ByRegion(geoLocInfo.Region)

	log.Printf("Active alerts in your region: today (%v), future (%v)\n", todayAlertsByRegion.Count(), futureAlertsByRegion.Count())
	if todayAlertsByRegion.Count() > 0 {
		for i, a := range todayAlertsByRegion {
			log.Printf("Today alert #%d: %s (%s)\n", i+1, a.Description, a.Severity)
		}
	}

	if futureAlertsByRegion.Count() > 0 {
		for i, a := range futureAlertsByRegion {
			log.Printf("Future alert #%d: %s (%s)\n", i+1, a.Description, a.Severity)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
