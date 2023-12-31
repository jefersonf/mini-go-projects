package inmet

type weatherAlerts []Alert

type ActiveAlerts struct {
	Today  weatherAlerts `json:"hoje"`
	Future weatherAlerts `json:"futuro"`
}

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

func (wa *weatherAlerts) Count() int {
	return len(*wa)
}

func (aa *ActiveAlerts) GetAll() (weatherAlerts, weatherAlerts) {
	return aa.TodayAlerts(), aa.FutureAlerts()
}

func (aa *ActiveAlerts) TodayAlerts() weatherAlerts {
	today := make(weatherAlerts, len(aa.Today))
	copy(today, aa.Today)
	return today
}

func (aa *ActiveAlerts) FutureAlerts() weatherAlerts {
	future := make(weatherAlerts, len(aa.Future))
	copy(future, aa.Future)
	return future
}

func (aa *ActiveAlerts) Count() int {
	return aa.Today.Count() + aa.Future.Count()
}
