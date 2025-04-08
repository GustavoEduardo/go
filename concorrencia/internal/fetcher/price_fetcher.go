package fetcher

import (
	"math/rand"
	"sync"
	"time"
)

// Buscar preços de sites com RPAs (simular)

// <- canal vai receber dados

func FetchPrices(priceChannel chan<- float64) {
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
	close(priceChannel) // Evita deadlock. Avisa as GO routines que não existe mais valor no canal

}

func FetchProceFromMercadoLivre() float64 {
	time.Sleep(1 * time.Second)

	return rand.Float64() * 100
}

func FetchProceFromKabum() float64 {
	time.Sleep(3 * time.Second)

	return rand.Float64() * 100
}

func FetchProceFromAmericanas() float64 {
	time.Sleep(2 * time.Second)

	return rand.Float64() * 100
}
