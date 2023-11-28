package models

type Time struct {
	ID           int64   `json:"id"`
	Nome_Time    string  `json:"nome_time"`
	Estado       string  `json:"estado"`
	Tecnico      string  `json:"tecnico"`
	Valor_Elenco float64 `json:"valor_elenco"`
}
