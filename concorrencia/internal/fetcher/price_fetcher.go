package fetcher

import (
	"math/rand"
	"time"
)

// Buscar preços de sites com RPAs (simular)

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
