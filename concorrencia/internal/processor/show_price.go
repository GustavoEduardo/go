package processor

import (
	"buscador/internal/models"
	"fmt"
)

// <- canal vai receber dados
// -> valor do canal será lido

func ShowPriceAVG(priceChannel <-chan models.PriceDetail, done chan<- bool) {

	var totalPrice float64
	countPrice := 0.0

	// O valor recebido em um channer só pode ser lido uma vez!
	for priceDetail := range priceChannel {
		totalPrice += priceDetail.Value
		countPrice++
		avgPrive := totalPrice / countPrice
		fmt.Printf("[%s] - Preço %s: R$ %.2f | Média : R$ %.2f \n",
			priceDetail.Timetamp.Format("02/Jan/2006 15:04:05"),
			priceDetail.StoreName, priceDetail.Value, avgPrive)
	}

	// https://go.dev/src/time/format.go

	done <- true

}
