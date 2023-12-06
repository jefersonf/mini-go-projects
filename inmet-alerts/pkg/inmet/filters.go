package inmet

import (
	"strings"
)

func (wa *weatherAlerts) ByRegion(region string) weatherAlerts {
	alerts := make(weatherAlerts, 0)
	for _, alert := range *wa {
		if strings.Contains(alert.States, region) {
			alerts = append(alerts, alert)
		}
	}
	return alerts
}
