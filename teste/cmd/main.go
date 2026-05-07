package main

import (
	"os"
	"teste/cmd/basics"
	"teste/internal/data"
	"teste/internal/handler"

	"github.com/gin-gonic/gin"

	nativo "fmt" // apelido
	_ "time"     // para importar sem usar
)

// exemplo de função
func LerArquivo() (*os.File, error) {

	arq, err := os.Open("")
	if err != nil {
		nativo.Println("Erro ao abrir o arquivo... ", err)
		// return nil, err
		panic("Erro que para a aplicação")
	}

	return arq, nil

}

func init() {

	// O go procura as funções init e main.
	// Todos as init são executadas antes.
	// Boa para checagens antes de rodar o main, instanciar serviços, etc.

}

func main() {

	basics.NumeroSecreto()

	data.LoadAtendimentos()

	router := gin.Default() // : pois estou iniciando a variável

	router.GET("/atendimento", handler.Get)
	router.GET("/atendimento/:id", handler.GetById)
	router.POST("/atendimento", handler.New)
	router.PUT("/atendimento/:id", handler.Update)
	router.DELETE("/atendimento/:id", handler.Remove)

	router.POST("/servicos", handler.Novo)

	router.Run()

}
