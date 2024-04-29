package types

import "time"

type CogSku struct {
	Price []int       `json:"price"`
	Date  []time.Time `json:"date"`
}
