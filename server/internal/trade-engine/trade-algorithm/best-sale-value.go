package trade_algorithm

import (
	"github.com/adriein/tibia-mkt/internal/trade-engine"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"sort"
)

type KeyValue struct {
	Key   int
	Value int
}

type BestSellValue struct {
	config *trade_engine.TradeEngineConfig
	prob   *service.ProbHelper
}

type SellOfferFrequency struct {
	Price       int     `json:"price"`
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

	offerFrequencyMap := make(map[int]int)
	prices := make([]int, len(cogs))

	for i := 0; i < len(cogs); i++ {
		prices[i] = cogs[i].SellPrice
	}

	historicAverage := int(bsv.prob.Mean(prices))

	for i := 0; i < len(cogs); i++ {
		appearance := offerFrequencyMap[cogs[i].SellPrice]

		if appearance > 0 {
			offerFrequencyMap[cogs[i].SellPrice] = appearance + 1

			continue
		}

		offerFrequencyMap[cogs[i].SellPrice] = 1
	}

	sortedKeyValue := bsv.sortFrequencyMap(offerFrequencyMap)[0:4]

	for _, keyValue := range sortedKeyValue {
		frequency := float64(keyValue.Value) / float64(len(cogs))

		frequencyResults = append(
			frequencyResults,
			SellOfferFrequency{Price: keyValue.Key, Occurrences: keyValue.Value, Frequency: frequency},
		)
	}

	result := BestSellValueResult{
		Mean:               historicAverage,
		StdDeviation:       bsv.prob.StdDeviation(prices),
		SellOfferFrequency: frequencyResults,
	}

	return result, nil
}

func (bsv *BestSellValue) sortFrequencyMap(offerFrequencyMap map[int]int) []KeyValue {
	var result []KeyValue
	for k, v := range offerFrequencyMap {
		result = append(result, KeyValue{Key: k, Value: v})
	}

	sort.Slice(result, func(i int, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}
