package handler

import (
	"net/http"
	"teste/internal/data"
	"teste/internal/models"
	"teste/internal/service"

	"github.com/gin-gonic/gin"
)

func Novo(c *gin.Context) {

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
