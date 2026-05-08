package main

// Funções

import (
	"fmt"
	"os"
)

// Letra maúscula compartilha a função com todo o app
// se tudo for a mesma tipagem, nao precisa tipar oque vem antes do tipo
// Ex (a, b int, c, d float) a = int e c = float

// vai devolver um int, por isso é uma func
// se não o GO entende como uma procedure
func Somar(a, b int) int {
	return a + b
}

// Vais rornar um inteiro mesmo que o resultado for float

func Dividir(a, b int) int {
	if b == 0 {
		// 	Sem o tratamento da panic: runtime error: integer divide by zero
		return 0
	}
	return a / b
}

// Retorna File ou nill
func LerArquivoSimples(caminho string) *os.File {

	arquivo, err := os.Open(caminho)

	if err != nil {
		fmt.Println("Erro ao abrir arquivo", err)
		return nil
	}

	return arquivo
}

/*
mesma funcionalidade de LerArquivoSimples, mas...
Retorna mais de um valor, sempre (arq, err, tamanho)
Nomeando os retornos e já aproveitando na atribuição
E com comentário colado, vira uma doc ao passar o mouse em cima da função!
Pode ser usado futuramento para gerar uma doc com o comando go doc ...
*/
func LerArquivo(caminho string) (arq *os.File, err error, tamanho int) {

	arq, err = os.Open(caminho)

	if err != nil {
		return nil, err, 0
	}

	// como o nome foi definido no retorno, não precisa aqui
	return
}

func main() {

	defer func() {
		fmt.Println("Func anônima. Por último por conta do defer")
	}()

	fmt.Println("Olá ------------------")
	fmt.Println(Somar(5, 10))
	fmt.Println(Dividir(45, 0))

	arq, err, _ := LerArquivo("arquivo.csv")
	// ignorei a variável tamanho

	if err != nil {
		fmt.Println("Erro ao abrir arquivo", err)
		return // Para parar o fluxo
	}

	defer arq.Close()
	// defer executa ao terminar a função (no caso a main)
	// Close executa antes da func anônima pq defer executa de baixo para cima

	// GO já tem o garbage collector que limpa a memória com o programa aberto
	// mas é bom impar a memória para não precisar esperar ele limpar sozinho
	// Ex. leio um arquivo grande e guardo em memória, fecho o arquivo, e só depois de editar o conteúdo. eu devolvo para o arquivo (abro de novo?)
	// Ex. tilizar o defer para fechar conexão.

	nome := "Gustavo"

	defer func(nomeRecebido string) {
		fmt.Println(nomeRecebido + " | " + nome)
	}(nome) // mesmo com defer, já gravou Gustavo, não adianta mudar

	nome = "Lima"

}

// defer é util para log, mandar email, etc
// recovery da aplicação também, ex:
// app caiu, e tento subir de novo, ou mando um email.
