package presenter

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/types"
	"time"
)

type CogAverage struct {
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
	ReferenceLine CogAverage   `json:"referenceLine"`
}

type DetailResponse struct {
	Wiki                  string                        `json:"wiki"`
	Creatures             []types.CreatureKillStatistic `json:"creatures"`
	SellOfferHistoricData []types.DataSnapshot          `json:"sellOfferHistoricData"`
	SellOfferProbability  SellOfferProbability          `json:"sellOfferProbability"`
	Cog                   []types.GoodResponse          `json:"cog"`
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

	cogSkuList := input.Cog

	var (
		buyOfferTotal        int
		sellOfferTotal       int
		cogSkuResponseList   []types.GoodResponse
		lowestSellPrice      types.Tick
		highestSellPrice     types.Tick
		lowestBuyPrice       types.Tick
		highestBuyPrice      types.Tick
		sellOfferYAxisDomain []types.Tick
		sellOfferXAxisDomain []string
		buyOfferYAxisDomain  []types.Tick
		buyOfferXAxisDomain  []string
	)

	highestSellPrice = types.Tick{Price: cogSkuList[0].SellPrice, Date: cogSkuList[0].Date.Format(time.DateOnly)}
	lowestSellPrice = types.Tick{Price: cogSkuList[0].SellPrice, Date: cogSkuList[0].Date.Format(time.DateOnly)}
	lowestBuyPrice = types.Tick{Price: cogSkuList[0].BuyPrice, Date: cogSkuList[0].Date.Format(time.DateOnly)}
	highestBuyPrice = types.Tick{Price: cogSkuList[0].BuyPrice, Date: cogSkuList[0].Date.Format(time.DateOnly)}

	for _, cogSku := range cogSkuList {
		buyOfferTotal = buyOfferTotal + cogSku.BuyPrice
		sellOfferTotal = sellOfferTotal + cogSku.SellPrice

		if highestSellPrice.Price < cogSku.SellPrice {
			highestSellPrice.Price = cogSku.SellPrice
			highestSellPrice.Date = cogSku.Date.Format(time.DateOnly)
		}

		if lowestSellPrice.Price > cogSku.SellPrice {
			lowestSellPrice.Price = cogSku.SellPrice
			lowestSellPrice.Date = cogSku.Date.Format(time.DateOnly)
		}

		if lowestBuyPrice.Price > cogSku.BuyPrice {
			lowestBuyPrice.Price = cogSku.BuyPrice
			lowestBuyPrice.Date = cogSku.Date.Format(time.DateOnly)
		}

		if highestBuyPrice.Price < cogSku.BuyPrice {
			highestBuyPrice.Price = cogSku.BuyPrice
			highestBuyPrice.Date = cogSku.Date.Format(time.DateOnly)
		}

		cogSkuResponseList = append(cogSkuResponseList, types.GoodResponse{
			BuyOffer:  cogSku.BuyPrice,
			SellOffer: cogSku.SellPrice,
			Date:      cogSku.Date.Format(time.DateOnly),
			World:     cogSku.World,
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

	result := DetailResponse{
		Wiki:                  input.Wiki,
		Creatures:             creatures,
		SellOfferHistoricData: input.SellOfferHistoricData,
		SellOfferProbability:  probability,
		Cog:                   cogSkuResponseList,
		SellOfferChart: DetailChartMetadataResponse{
			YAxisTick: sellOfferYAxisDomain,
			XAxisTick: sellOfferXAxisDomain,
			ReferenceLine: CogAverage{
				OfferType: constants.SellOfferType,
				Average:   p.calculateAverage(sellOfferTotal, len(cogSkuList)),
			},
		},
		BuyOfferChart: DetailChartMetadataResponse{
			YAxisTick: buyOfferYAxisDomain,
			XAxisTick: buyOfferXAxisDomain,
			ReferenceLine: CogAverage{
				OfferType: constants.BuyOfferType,
				Average:   p.calculateAverage(buyOfferTotal, len(cogSkuList)),
			},
		},
	}

	response := types.ServerResponse{
		Ok:   true,
		Data: result,
	}

	return response, nil
}

func (p *DetailPresenter) calculateAverage(totalSumCog int, totalNumCog int) int {
	if totalNumCog == 0 {
		return 0
	}

	return totalSumCog / totalNumCog
}
