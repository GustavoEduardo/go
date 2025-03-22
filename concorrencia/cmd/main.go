package main

import (
	"buscador/internal/fetcher"
	"fmt"
	"sync"
	"time"
)

func main() {

	start := time.Now()

	var priceMl, priceKb, priceAm float64
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		priceMl = fetcher.FetchProceFromMercadoLivre()
	}()

	go func() {
		defer wg.Done()
		priceKb = fetcher.FetchProceFromKabum()
	}()

	go func() {
		defer wg.Done()
		priceAm = fetcher.FetchProceFromAmericanas()
	}()

	wg.Wait()

	fmt.Printf("\nR$ %.2f\n", priceMl)
	fmt.Printf("R$ %.2f\n", priceKb)
	fmt.Printf("R$ %.2f\n", priceAm)

	fmt.Printf("\nTempo execução: %s\n", time.Since(start))

}
