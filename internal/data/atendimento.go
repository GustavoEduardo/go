package data

import (
	"encoding/json"
	"fmt"
	"os"
	"teste/internal/models"
)

var Atendimentos []models.Atendimento

func LoadAtendimentos() {

	file, err := os.Open("dados/atendimento.json")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo", err)
		return
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&Atendimentos); err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}
}

func SaveAtendimento() {

	file, err := os.Create("dados/atendimento.json")

	if err != nil {
		fmt.Println("Erro ao criar arquivo", err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(Atendimentos); err != nil {
		fmt.Println("Error encoding JSON: ", err)
	}

}
