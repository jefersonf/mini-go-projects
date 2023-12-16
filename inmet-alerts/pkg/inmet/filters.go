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

func (wa *weatherAlerts) BySeverity(sev string) weatherAlerts {
	alerts := make(weatherAlerts, 0)
	for _, alert := range *wa {
		if strings.Compare(alert.States, sev) == 0 {
			alerts = append(alerts, alert)
		}
	}
	return alerts
}
