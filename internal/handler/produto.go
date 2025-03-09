package handler

import (
	"net/http"
	"strconv"
	"teste/internal/data"
	"teste/internal/models"
	"teste/internal/service"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {

	c.JSON(200, gin.H{
		"produtos": data.Produtos,
	})

}

func GetById(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})

		return

	}

	for _, p := range data.Produtos {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Produto não encontrado!",
	})

}

func New(c *gin.Context) {

	var novoProduto models.Produto

	if err := c.ShouldBindJSON(&novoProduto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	if err := service.ValidateBody(&novoProduto); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	novoProduto.ID = len(data.Produtos) + 1

	data.Produtos = append(data.Produtos, novoProduto)

	data.SaveProduto()

	c.JSON(http.StatusCreated, gin.H{
		"menssagem": "Inserido com sucesso!",
		"produto":   novoProduto,
	})

}

func Update(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var updatedProduto models.Produto

	if err := c.ShouldBindJSON(&updatedProduto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := service.ValidateBody(&updatedProduto); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	for i, p := range data.Produtos {
		if p.ID == id {
			data.Produtos[i] = updatedProduto
			data.Produtos[i].ID = id
			data.SaveProduto()
			c.JSON(http.StatusCreated, gin.H{"message": "Editado com sucesso!", "data": data.Produtos[i]})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Produto não encontrado!"})

}

func SoftDelete(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	for i, p := range data.Produtos {
		if p.ID == id {
			data.Produtos = append(data.Produtos[:i], data.Produtos[i+1:]...) //até o indice i, do indice i+1 e continuar até o fim (:)
			data.SaveProduto()
			c.JSON(http.StatusOK, gin.H{"message": "Removido com sucesso!"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Produto não encontrado!"})

}

// Ver mais https://pkg.go.dev/net/http
