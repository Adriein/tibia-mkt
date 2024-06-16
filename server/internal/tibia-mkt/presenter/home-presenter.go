package presenter

import (
	"encoding/json"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/types"
	"time"
)

type HomeResponseCogSku struct {
	BuyPrice  int    `json:"buyPrice"`
	SellPrice int    `json:"sellPrice"`
	Date      string `json:"date"`
	World     string `json:"world"`
}

type ChartMetadata struct {
	XAxisTick []string `json:"xAxisTick"`
	YAxisTick []int    `json:"yAxisTick"`
}

type HomeResponse struct {
	Cogs  []HomeResponseCogSku `json:"cogs"`
	Chart ChartMetadata        `json:"chartMetadata"`
}

type HomePresenter struct{}

func NewHomePresenter() *HomePresenter {
	return &HomePresenter{}
}

func (p *HomePresenter) Format(data any) ([]byte, error) {
	cogSkuList, ok := data.([]types.CogSku)

	if !ok {
		return nil, types.ApiError{
			Msg:      "Assertion failed, data is not of type CogSku",
			Function: "Format",
			File:     "home-presenter.go",
		}
	}

	var (
		homeResponseList []HomeResponseCogSku
		highestSellPrice = cogSkuList[0].SellPrice
		lowestBuyPrice   = cogSkuList[0].BuyPrice
		yAxisDomain      []int
		xAxisDomain      []string
	)

	for _, cogSku := range cogSkuList {
		if highestSellPrice < cogSku.SellPrice {
			highestSellPrice = cogSku.SellPrice
		}

		if lowestBuyPrice > cogSku.SellPrice {
			lowestBuyPrice = cogSku.BuyPrice
		}

		homeResponseList = append(homeResponseList, HomeResponseCogSku{
			BuyPrice:  cogSku.BuyPrice,
			SellPrice: cogSku.SellPrice,
			Date:      cogSku.Date.Format(time.DateOnly),
			World:     cogSku.World,
		})
	}

	yAxisDomain = append(yAxisDomain, lowestBuyPrice, highestSellPrice)

	xAxisDomain = append(
		xAxisDomain,
		constants.Day1,
		constants.Day10,
		constants.Day20,
		constants.Day30,
		constants.Day31,
	)

	response := &types.ServerResponse{
		Ok: true,
		Data: HomeResponse{
			Cogs: homeResponseList,
			Chart: ChartMetadata{
				YAxisTick: yAxisDomain,
				XAxisTick: xAxisDomain,
			},
		},
	}

	bytes, jsonErr := json.Marshal(response)

	if jsonErr != nil {
		return nil, types.ApiError{
			Msg:      jsonErr.Error(),
			Function: "Format",
			File:     "home-presenter.go",
		}
	}

	return bytes, nil
}
