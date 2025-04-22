package main

import (
	"estoque/internal/models"
	"estoque/internal/services"
	"fmt"
)

func main() {
	fmt.Println("Sistema de stoque")

	estoque := services.NewEstoque()

	itens := []models.Item{
		{ID: "1", Name: "Produto 1", Quanity: 5, Price: 10.90},
		{ID: "2", Name: "Produto 2", Quanity: 10, Price: 15},
		{ID: "3", Name: "Produto 3", Quanity: 15, Price: 20.90},
	}

	for _, item := range itens {

		err := estoque.AddItem(item, "123")

		if err != nil {
			fmt.Println(err)
		}

	}

	items, err := estoque.ListItems()
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println(items)

	logs, err := estoque.ListLogs()
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("\nLogs: ", logs)

	itemsFound, err := services.Find(items, func(item models.Item) bool {
		return item.Price >= 15
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nItens encontrados: ", itemsFound)

	fornecedor := services.Fornecedor{
		CNPJ:    "123456",
		Contato: "11999999999",
		Cidade:  "São Paulo",
	}

	fmt.Println(fornecedor.GetInfo())
	fmt.Println(fornecedor.VerificarDisponibilidade(20, 22))

}
