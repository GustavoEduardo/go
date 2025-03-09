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

	data.LoadProdutos()

	router := gin.Default()

	router.GET("/produto", handler.Get)
	router.GET("/produto/:id", handler.GetById)
	router.POST("/produto", handler.New)
	router.PUT("/produto/:id", handler.Update)
	router.DELETE("/produto/:id", handler.SoftDelete)

	router.Run()

}
