package presenter

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/helper"
	"github.com/adriein/tibia-mkt/pkg/types"
	"time"
)

type GoodAverage struct {
	OfferType string `json:"offerType"`
	Average   int    `json:"average"`
}

type SellOfferProbability struct {
	Mean               int                        `json:"mean"`
	StdDeviation       float64                    `json:"stdDeviation"`
	SellOfferFrequency []types.SellOfferFrequency `json:"sellOfferFrequency"`
}

type DetailChartMetadataResponse struct {
	XAxisTick     []string     `json:"xAxisTick"`
	YAxisTick     []types.Tick `json:"yAxisTick"`
	ReferenceLine GoodAverage  `json:"referenceLine"`
}

type DetailResponse struct {
	Wiki                  string                        `json:"wiki"`
	Creatures             []types.CreatureKillStatistic `json:"creatures"`
	SellOfferHistoricData []types.DataSnapshot          `json:"sellOfferHistoricData"`
	SellOfferProbability  SellOfferProbability          `json:"sellOfferProbability"`
	Cogs                  []types.GoodRecordResponse    `json:"cogs"`
	SellOfferChart        DetailChartMetadataResponse   `json:"sellOfferChart"`
	BuyOfferChart         DetailChartMetadataResponse   `json:"buyOfferChart"`
}

type DetailPresenter struct{}

func NewDetailPresenter() *DetailPresenter {
	return &DetailPresenter{}
}

func (p *DetailPresenter) Format(data any) (types.ServerResponse, error) {
	input, ok := data.(types.Detail)

	if !ok {
		return types.ServerResponse{}, types.ApiError{
			Msg:      "Assertion failed, data is not an array of GoodRecord",
			Function: "Format",
			File:     "detail-presenter.go",
		}
	}

	goodSkuList := input.GoodRecord

	var (
		buyOfferTotal        int
		sellOfferTotal       int
		goodSkuResponseList  []types.GoodRecordResponse
		lowestSellPrice      types.Tick
		highestSellPrice     types.Tick
		lowestBuyPrice       types.Tick
		highestBuyPrice      types.Tick
		sellOfferYAxisDomain []types.Tick
		sellOfferXAxisDomain []string
		buyOfferYAxisDomain  []types.Tick
		buyOfferXAxisDomain  []string
	)

	highestSellPrice = types.Tick{Price: goodSkuList[0].SellPrice, Date: goodSkuList[0].Date.Format(time.DateOnly)}
	lowestSellPrice = types.Tick{Price: goodSkuList[0].SellPrice, Date: goodSkuList[0].Date.Format(time.DateOnly)}
	lowestBuyPrice = types.Tick{Price: goodSkuList[0].BuyPrice, Date: goodSkuList[0].Date.Format(time.DateOnly)}
	highestBuyPrice = types.Tick{Price: goodSkuList[0].BuyPrice, Date: goodSkuList[0].Date.Format(time.DateOnly)}

	for _, goodSku := range goodSkuList {
		buyOfferTotal = buyOfferTotal + goodSku.BuyPrice
		sellOfferTotal = sellOfferTotal + goodSku.SellPrice

		if highestSellPrice.Price < goodSku.SellPrice {
			highestSellPrice.Price = goodSku.SellPrice
			highestSellPrice.Date = goodSku.Date.Format(time.DateOnly)
		}

		if lowestSellPrice.Price > goodSku.SellPrice {
			lowestSellPrice.Price = goodSku.SellPrice
			lowestSellPrice.Date = goodSku.Date.Format(time.DateOnly)
		}

		if lowestBuyPrice.Price > goodSku.BuyPrice {
			lowestBuyPrice.Price = goodSku.BuyPrice
			lowestBuyPrice.Date = goodSku.Date.Format(time.DateOnly)
		}

		if highestBuyPrice.Price < goodSku.BuyPrice {
			highestBuyPrice.Price = goodSku.BuyPrice
			highestBuyPrice.Date = goodSku.Date.Format(time.DateOnly)
		}

		goodSkuResponseList = append(goodSkuResponseList, types.GoodRecordResponse{
			BuyOffer:  goodSku.BuyPrice,
			SellOffer: goodSku.SellPrice,
			Date:      goodSku.Date.Format(time.DateOnly),
			World:     goodSku.World,
		})
	}

	sellOfferYAxisDomain = append(sellOfferYAxisDomain, lowestSellPrice, highestSellPrice)
	buyOfferYAxisDomain = append(buyOfferYAxisDomain, lowestBuyPrice, highestBuyPrice)

	buyOfferXAxisDomain = append(
		buyOfferXAxisDomain,
		constants.Day1,
		constants.Day10,
		constants.Day20,
		constants.Day30,
		constants.Day31,
	)

	sellOfferXAxisDomain = append(
		sellOfferXAxisDomain,
		constants.Day1,
		constants.Day10,
		constants.Day20,
		constants.Day30,
		constants.Day31,
	)

	creatures := input.Creatures

	if input.Creatures == nil || len(input.Creatures) == 0 {
		creatures = make([]types.CreatureKillStatistic, 0)
	}

	probability := SellOfferProbability{
		StdDeviation:       input.StdDeviation,
		Mean:               input.SellPriceMean,
		SellOfferFrequency: input.SellOfferFrequency,
	}

	helper.Reverse[types.DataSnapshot](input.SellOfferHistoricData)

	result := DetailResponse{
		Wiki:                  input.Wiki,
		Creatures:             creatures,
		SellOfferHistoricData: input.SellOfferHistoricData,
		SellOfferProbability:  probability,
		Cogs:                  goodSkuResponseList,
		SellOfferChart: DetailChartMetadataResponse{
			YAxisTick: sellOfferYAxisDomain,
			XAxisTick: sellOfferXAxisDomain,
			ReferenceLine: GoodAverage{
				OfferType: constants.SellOfferType,
				Average:   p.calculateAverage(sellOfferTotal, len(goodSkuList)),
			},
		},
		BuyOfferChart: DetailChartMetadataResponse{
			YAxisTick: buyOfferYAxisDomain,
			XAxisTick: buyOfferXAxisDomain,
			ReferenceLine: GoodAverage{
				OfferType: constants.BuyOfferType,
				Average:   p.calculateAverage(buyOfferTotal, len(goodSkuList)),
			},
		},
	}

	response := types.ServerResponse{
		Ok:   true,
		Data: result,
	}

	return response, nil
}

func (p *DetailPresenter) calculateAverage(totalSumGood int, totalNumGood int) int {
	if totalNumGood == 0 {
		return 0
	}

	return totalSumGood / totalNumGood
}
