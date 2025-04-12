package main

import (
	"estoque/internal/models"
	"fmt"
)

func main() {
	fmt.Println("Sistema de stoque")

	item1 := models.Item{
		ID:      1,
		Name:    "Produto 1",
		Quanity: 10,
		Price:   19.90,
	}

	fmt.Println(item1.Info())
}
