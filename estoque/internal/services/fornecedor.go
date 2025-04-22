package services

import "fmt"

// implementação de interfaces  em GO é implicita

type InfoService interface {
	GetInfo() string
}

type DisponibilidadeService interface {
	VerificarDisponibilidade(qtdSolicitada int, qtdDisponivel int) bool
}

// Interface encadeada
type FornecedorService interface {
	InfoService
	DisponibilidadeService
}

type Fornecedor struct {
	CNPJ    string
	Contato string
	Cidade  string
}

func (f Fornecedor) GetInfo() string {
	return fmt.Sprintf("\nCNPJ: %s | Contato: %s", f.CNPJ, f.Contato)
}

func (f Fornecedor) VerificarDisponibilidade(qtdSolicitada int, qtdDisponivel int) bool {
	return qtdSolicitada <= qtdDisponivel
}
