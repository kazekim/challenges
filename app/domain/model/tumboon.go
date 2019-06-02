package model

import "time"

type Tumboon struct {
	Name string
	AmountSubUnits int64 `csv:"AmountSubunits"`
	CCNumber string
	CVV string
	ExpMonth time.Month
	ExpYear int
}


