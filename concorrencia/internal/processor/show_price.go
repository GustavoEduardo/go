package processor

import "fmt"

// <- canal vai receber dados
// -> valor do canal será lido

func ShowPriceAVG(priceChannel <-chan float64, done chan<- bool) {

	var totalPrice float64
	countPrice := 0.0

	// O valor recebido em um channer só pode ser lido uma vez!
	for price := range priceChannel {
		totalPrice += price
		countPrice++
		avgPrive := totalPrice / countPrice
		fmt.Printf("Preço recebido : R$ %.2f | Média : R$ %.2f \n", price, avgPrive)
	}

	done <- true

}
