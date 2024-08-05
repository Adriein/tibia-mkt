package trade_algorithm

import (
	"fmt"
	"github.com/adriein/tibia-mkt/internal/trade-engine"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"sort"
)

type KeyValue struct {
	Key   string
	Value int
}

type BestSellValue struct {
	config *trade_engine.TradeEngineConfig
	prob   *service.ProbHelper
}

type SellOfferFrequency struct {
	Range       string  `json:"range"`
	Occurrences int     `json:"occurrences"`
	Frequency   float64 `json:"frequency"`
}

type BestSellValueResult struct {
	Mean               int                  `json:"mean"`
	StdDeviation       float64              `json:"stdDeviation"`
	SellOfferFrequency []SellOfferFrequency `json:"sellOfferFrequency"`
}

func NewBestSellValueAlgorithm(config *trade_engine.TradeEngineConfig, prob *service.ProbHelper) *BestSellValue {
	return &BestSellValue{config: config, prob: prob}
}

func (bsv *BestSellValue) Apply(cogs []types.CogSku) (BestSellValueResult, error) {
	var (
		frequencyResults []SellOfferFrequency
	)

	priceRanges := make(map[string]int)
	prices := make([]int, len(cogs))

	for i := 0; i < len(cogs); i++ {
		prices[i] = cogs[i].SellPrice

		rangeStart := (cogs[i].SellPrice / 100) * 100
		rangeEnd := rangeStart + 99

		rangeStr := fmt.Sprintf("%d-%d", rangeStart, rangeEnd)

		priceRanges[rangeStr]++
	}

	sortedKeyValue := bsv.sortFrequencyMap(priceRanges)

	for _, keyValue := range sortedKeyValue {
		frequency := float64(keyValue.Value) / float64(len(cogs))

		frequencyResults = append(
			frequencyResults,
			SellOfferFrequency{Range: keyValue.Key, Occurrences: keyValue.Value, Frequency: frequency},
		)
	}

	result := BestSellValueResult{
		Mean:               int(bsv.prob.Mean(prices)),
		StdDeviation:       bsv.prob.StdDeviation(prices),
		SellOfferFrequency: frequencyResults,
	}

	return result, nil
}

func (bsv *BestSellValue) sortFrequencyMap(priceRanges map[string]int) []KeyValue {
	var result []KeyValue
	for k, v := range priceRanges {
		result = append(result, KeyValue{Key: k, Value: v})
	}

	sort.Slice(result, func(i int, j int) bool {
		return result[i].Key < result[j].Key
	})

	return result
}
