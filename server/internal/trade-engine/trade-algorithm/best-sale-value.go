package trade_algorithm

import (
	"github.com/adriein/tibia-mkt/internal/trade-engine"
	"github.com/adriein/tibia-mkt/pkg/types"
)

type BestSellValue struct {
	config trade_engine.TradeEngineConfig
}

func (bsv *BestSellValue) Apply(cogs []types.CogSku) error {
	var (
		totalCogSellPrice int
	)

	for i := 0; i < len(cogs); i++ {
		totalCogSellPrice = totalCogSellPrice + cogs[i].SellPrice
	}

	historicAverage := totalCogSellPrice / len(cogs)

	averageFee := historicAverage * (bsv.config.OfferCreationFeePercentage / 100)

	for i := 0; i < len(cogs); i++ {
		delta := historicAverage - cogs[i].SellPrice

		if delta > averageFee {

		}
	}

	return nil
}
