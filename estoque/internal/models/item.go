package models

import "fmt"

type Item struct {
	ID      string
	Name    string
	Quanity int
	Price   float64
}

// O método info é uma função vinculada da struct e precisa de um receptor antes do nome.
// O receptor é equivalente ao self ou this de outras linguagens
// Por convenção usar a primeira letra do tipo

func (i Item) Info() string {
	return fmt.Sprintf("ID %s | Descrição: %s | Quantidade: %d | Valor: %.2f ", i.ID, i.Name, i.Quanity, i.Price)
}
