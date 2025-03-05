package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"teste/models"

	"github.com/gin-gonic/gin"
)

var produtos []models.Produto

func main() {

	// go get = npm i

	// API com gim...

	loadProdutos() // para persistir os dados

	router := gin.Default()

	router.GET("/produto", getProdutos)
	router.GET("/produto/:id", getProdutoById)
	router.POST("/produto", novoProduto)

	router.Run()

}

func getProdutos(c *gin.Context) {

	c.JSON(200, gin.H{
		"produtos": produtos,
	})

}

func getProdutoById(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {

		c.JSON(400, gin.H{
			"erro": err.Error(),
		})

		return

	}

	for _, p := range produtos {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}

	c.JSON(404, gin.H{
		"message": "Produto não encontrado!",
	})

}

func novoProduto(c *gin.Context) {

	var novoProduto models.Produto

	if err := c.ShouldBindJSON(&novoProduto); err != nil {
		c.JSON(400, gin.H{
			"erro": "Erro ao criar produto: " + err.Error(),
		})
		return
	}

	novoProduto.ID = len(produtos) + 1

	produtos = append(produtos, novoProduto)

	saveProduto()

	c.JSON(201, gin.H{
		"menssagem": "Produto inserido com sucesso!",
		"produto":   novoProduto,
	})

}

// -------------------------------------- Arquivo JSON

func loadProdutos() {

	file, err := os.Open("dados/produto.json")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo", err)
		return
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&produtos); err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}

}

func saveProduto() {

	file, err := os.Create("dados/produto.json")

	if err != nil {
		fmt.Println("Erro ao criar arquivo", err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(produtos); err != nil {
		fmt.Println("Error encoding JSON: ", err)
	}

}
