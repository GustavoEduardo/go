package models

import "time"

type PriceDetail struct {
	StoreName string
	Value     float64
	Timetamp  time.Time
}
