package inmet

type ActiveAlerts struct {
	Today  []Alert `json:"hoje"`
	Future []Alert `json:"futuro"`
}

type Alert struct {
	ID                uint     `json:"id"`
	AlertID           uint     `json:"id_aviso"`
	SequenceID        uint     `json:"id_sequencia"`
	SevereConditionID uint     `json:"id_condicao_severa"`
	IconID            uint     `json:"id_icone"`
	UserID            uint     `json:"id_usuario"`
	Code              string   `json:"codigo"`
	Reference         any      `json:"referencia,omitempty"`
	StartDate         string   `json:"data_inicio"`
	StartHour         string   `json:"hora_inicio"`
	EndDate           string   `json:"data_fim"`
	EndHour           string   `json:"hora_fim"`
	Poligon           string   `json:"poligono"`
	Cities            string   `json:"municipios"`
	Microregions      string   `json:"microrregioes"`
	Mesoregions       string   `json:"mesorregioes"`
	States            string   `json:"estados"`
	Regions           string   `json:"regioes"`
	Geocodes          string   `json:"geocodes"`
	Modified          bool     `json:"alterado"`
	Closed            bool     `json:"encerrado"`
	CreatedAt         string   `json:"created_at"`
	UpdatedAt         string   `json:"updated_at"`
	Start             string   `json:"inicio"`
	End               string   `json:"fim"`
	Icon              string   `json:"icone"`
	Description       string   `json:"descricao"`
	AlertColor        string   `json:"aviso_cor"`
	SeverityID        uint     `json:"id_severidade"`
	Severity          string   `json:"severidade"`
	Risks             []string `json:"riscos"`
	Instructions      []string `json:"instrucoes"`
}
