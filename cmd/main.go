package main

import (
	"teste/cmd/basics"
	"teste/internal/data"
	"teste/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	// go get = npm i

	// API com gim:

	basics.NumeroSecreto()

	data.LoadAtendimentos()

	router := gin.Default()

	router.GET("/atendimento", handler.Get)
	router.GET("/atendimento/:id", handler.GetById)
	router.POST("/atendimento", handler.New)
	router.PUT("/atendimento/:id", handler.Update)
	router.DELETE("/atendimento/:id", handler.SoftDelete)

	router.Run()

}
