package repository

import (
	"github.com/adriein/exori-vis-trade/pkg/types"
	"time"
)

type CsvSecuraCogRepository struct {
}

func NewCsvSecuraCogRepository() *CsvSecuraCogRepository {
	return &CsvSecuraCogRepository{}
}

func (c *CsvSecuraCogRepository) Find(criteria types.Criteria) []types.CogSku {
	var result []types.CogSku
	var dates []time.Time
	var prices []int

	dates = append(dates, time.Now())
	prices = append(prices, 1)

	result = append(result, types.CogSku{
		Date:  dates,
		Price: prices,
	})

	return result
}
