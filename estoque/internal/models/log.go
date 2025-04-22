package models

import "time"

type Log struct {
	Timestamp time.Time
	Action    string
	UserId    string
	ItemId    string
	Quantity  int
	Reason    string
}
