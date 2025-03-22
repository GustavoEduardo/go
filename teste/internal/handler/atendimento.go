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

	if data.Atendimentos == nil {
		c.JSON(http.StatusOK, gin.H{
			"atendimentos": []models.Atendimento{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"atendimentos": data.Atendimentos,
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

	for _, p := range data.Atendimentos {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Atendimento não encontrado!",
	})

}

func New(c *gin.Context) {

	var novoAtendimento models.Atendimento

	if err := c.ShouldBindJSON(&novoAtendimento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	if err := service.ValidateBody(&novoAtendimento); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	novoAtendimento.ID = len(data.Atendimentos) + 1

	data.Atendimentos = append(data.Atendimentos, novoAtendimento)

	data.SaveAtendimento()

	c.JSON(http.StatusCreated, gin.H{
		"menssagem":   "Inserido com sucesso!",
		"atendimento": novoAtendimento,
	})

}

func Update(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var updatedAtendimento models.Atendimento

	if err := c.ShouldBindJSON(&updatedAtendimento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := service.ValidateBody(&updatedAtendimento); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	for i, p := range data.Atendimentos {
		if p.ID == id {
			data.Atendimentos[i] = updatedAtendimento
			data.Atendimentos[i].ID = id
			data.SaveAtendimento()
			c.JSON(http.StatusCreated, gin.H{"message": "Editado com sucesso!", "data": data.Atendimentos[i]})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Atendimento não encontrado!"})

}

func SoftDelete(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	for i, p := range data.Atendimentos {
		if p.ID == id {
			data.Atendimentos = append(data.Atendimentos[:i], data.Atendimentos[i+1:]...) //até o indice i, do indice i+1 e continuar até o fim (:)
			data.SaveAtendimento()
			c.JSON(http.StatusOK, gin.H{"message": "Removido com sucesso!"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Atendimento não encontrado!"})

}

// Ver mais https://pkg.go.dev/net/http
