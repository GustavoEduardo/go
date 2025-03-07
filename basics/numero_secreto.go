package basics

import (
	"fmt"
	"math/rand"
	"strconv"
)

// ------------------------------------------------ Jogo número secreto

func numeroSecreto() {
	fmt.Println("Boas vindas ao jogo do número secreto")
	numeroMaximo := 1000
	// rand.Seed(time.Now().UnixNano()) deprecated (Go 1.5 > já prepara p gerador)
	// rand.New(rand.NewSource(time.Now().UnixNano())) < isso é feito por baixo dos panos
	numeroSecreto := rand.Intn(numeroMaximo) + 1
	fmt.Printf("Número secreto (para debug): %d\n", numeroSecreto)
	var chute int
	tentativas := 1 // mesma coisa que var tentativas int = 1

	for {
		fmt.Printf("Escolha um número entre 1 e %d: ", numeroMaximo)
		var input string
		fmt.Scanln(&input)
		chute, _ = strconv.Atoi(input) // basicamente, tudo retorna dois valores (result, err), e _ omite oque eu não quero usar

		if chute == numeroSecreto {
			break
		} else {
			if chute > numeroSecreto {
				fmt.Printf("O número secreto é menor que %d\n", chute)
			} else {
				fmt.Printf("O número secreto é maior que %d\n", chute)
			}
			tentativas++
		}
	}

	palavraTentativa := "tentativa"
	if tentativas > 1 {
		palavraTentativa = "tentativas"
	}

	fmt.Printf("\nIsso aí! Você descobriu o número secreto %d com %d %s.\n\n", numeroSecreto, tentativas, palavraTentativa)

	// ------------------------------------------------ Saída de dados
	// idade := 30
	// preco := 19.99
	// nome := "João"
	// status := true
	// numero := 255

	// fmt.Printf("Idade: %d anos \n", idade)
	// fmt.Printf("Preço: %.2f \n", preco)
	// fmt.Printf("Nome: %s \n", nome)
	// fmt.Printf("Status: %t \n", status)
	// fmt.Printf("Número em hexadecimal: %x \n", numero)

}
