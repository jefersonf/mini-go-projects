package inmet

type weatherAlerts []Alert

// ActiveAlerts describes the INMET summarized data.
type ActiveAlerts struct {
	Today  weatherAlerts `json:"hoje"`
	Future weatherAlerts `json:"futuro"`
}

// BaseAlert describes the basic properties of an alert to be listed.
type BaseAlert struct {
	SevereConditionID uint   `json:"id_condicao_severa"`
	StartDate         string `json:"data_inicio"`
	StartHour         string `json:"hora_inicio"`
	EndDate           string `json:"data_fim"`
	EndHour           string `json:"hora_fim"`
	Poligon           string `json:"poligono"`
	Cities            string `json:"municipios"`
	Microregions      string `json:"microrregioes"`
	Mesoregions       string `json:"mesorregioes"`
	States            string `json:"estados"`
	Regions           string `json:"regioes"`
	Geocodes          string `json:"geocodes"`
}

// Alert describes all the properties of an alert.
type Alert struct {
	BaseAlert
	ID           uint     `json:"id"`
	AlertID      uint     `json:"id_aviso"`
	SequenceID   uint     `json:"id_sequencia"`
	IconID       uint     `json:"id_icone"`
	UserID       uint     `json:"id_usuario"`
	Code         string   `json:"codigo"`
	Reference    any      `json:"referencia,omitempty"`
	Modified     bool     `json:"alterado"`
	Closed       bool     `json:"encerrado"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
	Start        string   `json:"inicio"`
	End          string   `json:"fim"`
	Icon         string   `json:"icone"`
	Description  string   `json:"descricao"`
	AlertColor   string   `json:"aviso_cor"`
	SeverityID   uint     `json:"id_severidade"`
	Severity     string   `json:"severidade"`
	Risks        []string `json:"riscos"`
	Instructions []string `json:"instrucoes"`
}

// Count returns the number of alerts of in list of alerts.
func (wa *weatherAlerts) Count() int {
	return len(*wa)
}

// GetAll returns two different lists of weather alerts, 
// the first is the today's list of alerts and the secon are the future alerts.
func (aa *ActiveAlerts) GetAll() (weatherAlerts, weatherAlerts) {
	return aa.TodayAlerts(), aa.FutureAlerts()
}

// TodayAlerts return the today's list of active alerts.
func (aa *ActiveAlerts) TodayAlerts() weatherAlerts {
	today := make(weatherAlerts, len(aa.Today))
	copy(today, aa.Today)
	return today
}

// FutureAlerts returns the future alerts.
func (aa *ActiveAlerts) FutureAlerts() weatherAlerts {
	future := make(weatherAlerts, len(aa.Future))
	copy(future, aa.Future)
	return future
}

// Count returns all active alerts. Both future and today's alerts.
func (aa *ActiveAlerts) Count() int {
	return aa.Today.Count() + aa.Future.Count()
}
