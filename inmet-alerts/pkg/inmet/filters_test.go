package inmet

import "testing"

func TestFilterByRegion(t *testing.T) {

	testcase := struct {
		alerts      ActiveAlerts
		region      string
		todayCount  int
		futureCount int
	}{
		alerts: ActiveAlerts{
			Today: weatherAlerts{
				{ID: 100, Description: "Chuvas Intensas", BaseAlert: BaseAlert{States: "Bahia,Paraíba,Pernambuco"}},
				{ID: 101, Description: "Tempestades", BaseAlert: BaseAlert{States: "Paraíba,Rio Grande do Norte"}},
			},
			Future: weatherAlerts{
				{ID: 102, Description: "Baixa Umidade", BaseAlert: BaseAlert{States: "São Paulo,Minas Gerais"}},
			},
		},
		region:      "Paraíba",
		todayCount:  2,
		futureCount: 0,
	}

	todayByRegion := testcase.alerts.Today.ByRegion(testcase.region)
	todayCount := len(todayByRegion)
	if todayCount != testcase.todayCount {
		t.Errorf("TodayByRegion got %v, want %v\n", todayCount, testcase.todayCount)
	}

	futureByRegion := testcase.alerts.Future.ByRegion(testcase.region)
	futureCount := len(futureByRegion)
	if futureCount != testcase.futureCount {
		t.Errorf("FutureByRegion got %v, want %v\n", futureByRegion, testcase.futureCount)
	}
}
