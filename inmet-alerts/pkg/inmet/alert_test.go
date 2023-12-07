package inmet

import (
	"strings"
	"testing"
)

func TestGetAll(t *testing.T) {
	testcases := []struct {
		alerts           ActiveAlerts
		TodayAlertsDesc  []string
		FutureAlertsDesc []string
	}{
		{
			alerts: ActiveAlerts{
				Today:  weatherAlerts{{ID: 100, Description: "Chuvas Intensas"}, {ID: 101, Description: "Tempestades"}},
				Future: weatherAlerts{{ID: 102, Description: "Baixa Umidade"}},
			},
			TodayAlertsDesc:  []string{"Chuvas Intensas", "Tempestades"},
			FutureAlertsDesc: []string{"Baixa Umidade"},
		},
	}

	for _, tc := range testcases {
		today, future := tc.alerts.GetAll()
		if today.Count() != len(tc.TodayAlertsDesc) {
			t.Errorf("TodayAlertsCount got %v, %v\n", today.Count(), len(tc.TodayAlertsDesc))
		}

		for i := range today {
			if strings.Compare(today[i].Description, tc.TodayAlertsDesc[i]) != 0 {
				t.Errorf("TodayAlertDesc got %v, %v\n", today.Count(), len(tc.TodayAlertsDesc))
			}
		}

		if future.Count() != len(tc.FutureAlertsDesc) {
			t.Errorf("FutureAlertsCount got %v, %v\n", future.Count(), len(tc.FutureAlertsDesc))
		}

		for i := range future {
			if strings.Compare(future[i].Description, tc.FutureAlertsDesc[i]) != 0 {
				t.Errorf("FutureAlertDesc got %v, %v\n", future.Count(), len(tc.FutureAlertsDesc))
			}
		}
	}
}

func TestCount_AllAlerts(t *testing.T) {
	sampleAlerts := ActiveAlerts{
		Today:  weatherAlerts{{ID: 100, Description: "Chuvas Intensas"}, {ID: 101, Description: "Tempestades"}},
		Future: weatherAlerts{{ID: 102, Description: "Baixa Umidade"}},
	}
	count := sampleAlerts.Count()
	if count != (sampleAlerts.Today.Count() + sampleAlerts.Future.Count()) {
		t.Errorf("got %v, want %v\n", sampleAlerts.Count(), 0)
	}
}

func TestCount_NoAlerts(t *testing.T) {
	sampleAlerts := make(weatherAlerts, 0)
	count := sampleAlerts.Count()
	if count != 0 {
		t.Errorf("got %v, want %v\n", sampleAlerts.Count(), 0)
	}
}
