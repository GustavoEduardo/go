package fetcher

import (
	"buscador/internal/models"
	"math/rand"
	"sync"
	"time"
)

// <- canal vai receber dados

func FetchPrices(priceChannel chan<- models.PriceDetail) {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		priceChannel <- FetchProceFromMercadoLivre()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchProceFromKabum()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchProceFromAmericanas()
	}()

	wg.Wait()
	close(priceChannel) // Evita deadlock. Avisa as GO routines que não existe mais valor para ser lido no canal

}

// Buscar preços de sites com RPAs (simular)

func FetchProceFromMercadoLivre() models.PriceDetail {
	time.Sleep(1 * time.Second)

	return models.PriceDetail{
		StoreName: "Mercado Livre",
		Value:     rand.Float64() * 100,
		Timetamp:  time.Now(),
	}
}

func FetchProceFromKabum() models.PriceDetail {
	time.Sleep(3 * time.Second)

	return models.PriceDetail{
		StoreName: "Kabum",
		Value:     rand.Float64() * 100,
		Timetamp:  time.Now(),
	}
}

func FetchProceFromAmericanas() models.PriceDetail {
	time.Sleep(2 * time.Second)

	return models.PriceDetail{
		StoreName: "Americanas",
		Value:     rand.Float64() * 100,
		Timetamp:  time.Now(),
	}
}
