package trade_algorithm

import (
	"github.com/adriein/tibia-mkt/internal/trade-engine"
	"github.com/adriein/tibia-mkt/pkg/types"
	"sort"
)

type KeyValue struct {
	Key   int
	Value int
}

type BestSellValue struct {
	config *trade_engine.TradeEngineConfig
}

type SellOfferFrequency struct {
	Price     int
	Frequency int
}

type BestSellValueResult struct {
	HistoricAveragePrice int
	SellOfferFrequency   []SellOfferFrequency
}

func NewBestSellValueAlgorithm(config *trade_engine.TradeEngineConfig) *BestSellValue {
	return &BestSellValue{config: config}
}

func (bsv *BestSellValue) Apply(cogs []types.CogSku) (BestSellValueResult, error) {
	var (
		totalCogSellPrice int
		frequencyResults  []SellOfferFrequency
	)

	offerFrequencyMap := make(map[int]int)

	for i := 0; i < len(cogs); i++ {
		totalCogSellPrice = totalCogSellPrice + cogs[i].SellPrice
	}

	historicAverage := totalCogSellPrice / len(cogs)

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
		frequency := keyValue.Value / len(cogs)

		frequencyResults = append(frequencyResults, SellOfferFrequency{Price: keyValue.Key, Frequency: frequency})
	}

	result := BestSellValueResult{
		HistoricAveragePrice: historicAverage,
		SellOfferFrequency:   frequencyResults,
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
