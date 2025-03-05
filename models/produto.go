package models

// ------------------------------------------------- struct

// atributos quei iniciam com letra minúscula são privados.
// para retornar como minpusculo usar a referência `json:"nomeDoCampo"`

type Produto struct {
	ID           int     `json:"id"`
	Descricao    string  `json:"descricao"`
	Valor        float64 `json:"valor"`
	campoPrivado string
}
