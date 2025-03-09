package data

import (
	"encoding/json"
	"fmt"
	"os"
	"teste/internal/models"
)

var Produtos []models.Produto

func LoadProdutos() {

	file, err := os.Open("dados/produto.json")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo", err)
		return
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&Produtos); err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}

}

func SaveProduto() {

	file, err := os.Create("dados/produto.json")

	if err != nil {
		fmt.Println("Erro ao criar arquivo", err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(Produtos); err != nil {
		fmt.Println("Error encoding JSON: ", err)
	}

}
