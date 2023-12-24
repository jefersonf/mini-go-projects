package inmet

import (
	"strings"
)

// ByRegion returns weather alerts filtered by a given region name.
func (wa *weatherAlerts) ByRegion(region string) weatherAlerts {
	alerts := make(weatherAlerts, 0)
	for _, alert := range *wa {
		if strings.Contains(alert.States, region) {
			alerts = append(alerts, alert)
		}
	}
	return alerts
}

// BySeverity returns weather alerts filtered by a given severity type.
func (wa *weatherAlerts) BySeverity(severity string) weatherAlerts {
	alerts := make(weatherAlerts, 0)
	for _, alert := range *wa {
		if strings.Compare(strings.ToLower(alert.States), strings.ToLower(severity)) == 0 {
			alerts = append(alerts, alert)
		}
	}
	return alerts
}
