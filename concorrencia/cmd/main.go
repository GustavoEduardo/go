package main

import (
	"buscador/internal/fetcher"
	"buscador/internal/models"
	"buscador/internal/processor"
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	priceChannel := make(chan models.PriceDetail)

	/*
		Canal com buffer. Controlo a qtd de leitura e escrita para uma possível economia de recursos.
		Pode umentar a complexidade.

		priceChannel := make(chan models.PriceDetail, 4)

		len -> Qtd de valores no buffer. Se entras um 5º, ele vai esperar um ser lido!
		cap -> fTamanhho maximo di buffer (4)
	*/

	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAVG(priceChannel, done)

	<-done

	fmt.Printf("\nTempo execução: %s\n", time.Since(start))

}
