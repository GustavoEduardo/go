package models

type Servico struct {
	Descricao      string  `json:"descricao"`
	Valor          float64 `json:"valor"`
	TempoEmMinutos int     `json:"tempo_em_minutos"`
}
