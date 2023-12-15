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

// BySeverity filter alerts by alert serevity name.
func (wa *weatherAlerts) BySeverity(severity string) weatherAlerts {
	alerts := make(weatherAlerts, 0)
	for _, alert := range *wa {
		if strings.Contains(strings.ToLower(alert.Severity), strings.ToLower(severity)) {
			alerts = append(alerts, alert)
		}
	}
	return alerts
}
