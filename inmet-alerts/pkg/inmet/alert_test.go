package inmet

import (
	"strings"
	"testing"
)

func TestGetAll(t *testing.T) {
	testcases := []struct {
		alerts           ActiveAlerts
		todayAlertsDesc  []string
		futureAlertsDesc []string
	}{
		{
			alerts: ActiveAlerts{
				Today:  weatherAlerts{{ID: 100, Description: "Chuvas Intensas"}, {ID: 101, Description: "Tempestades"}},
				Future: weatherAlerts{{ID: 102, Description: "Baixa Umidade"}},
			},
			todayAlertsDesc:  []string{"Chuvas Intensas", "Tempestades"},
			futureAlertsDesc: []string{"Baixa Umidade"},
		},
	}

	for _, tc := range testcases {
		today, future := tc.alerts.GetAll()
		if today.Count() != len(tc.todayAlertsDesc) {
			t.Errorf("TodayAlertsCount got %v, %v\n", today.Count(), len(tc.todayAlertsDesc))
		}

		for i := range today {
			if strings.Compare(today[i].Description, tc.todayAlertsDesc[i]) != 0 {
				t.Errorf("TodayAlertDesc got %v, %v\n", today.Count(), len(tc.todayAlertsDesc))
			}
		}

		if future.Count() != len(tc.futureAlertsDesc) {
			t.Errorf("FutureAlertsCount got %v, %v\n", future.Count(), len(tc.futureAlertsDesc))
		}

		for i := range future {
			if strings.Compare(future[i].Description, tc.futureAlertsDesc[i]) != 0 {
				t.Errorf("FutureAlertDesc got %v, %v\n", future.Count(), len(tc.futureAlertsDesc))
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
