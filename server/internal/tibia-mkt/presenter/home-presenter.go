package presenter

import (
	"encoding/json"
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

	var (
		homeResponseList []HomeResponseCogSku
		highestSellPrice = cogSkuList[0].SellPrice
		lowestSellPrice  = cogSkuList[0].SellPrice
		yAxisTick        []int
	)

	if !ok {
		return nil, types.ApiError{
			Msg:      "Assertion failed, data is not of type CogSku",
			Function: "Format",
			File:     "home-presenter.go",
		}
	}

	for _, cogSku := range cogSkuList {
		if highestSellPrice < cogSku.SellPrice {
			highestSellPrice = cogSku.SellPrice
		}

		if lowestSellPrice > cogSku.SellPrice {
			lowestSellPrice = cogSku.SellPrice
		}

		homeResponseList = append(homeResponseList, HomeResponseCogSku{
			BuyPrice:  cogSku.BuyPrice,
			SellPrice: cogSku.SellPrice,
			Date:      cogSku.Date.Format(time.DateOnly),
			World:     cogSku.World,
		})
	}

	yAxisTick = append(yAxisTick, lowestSellPrice, highestSellPrice)

	response := &types.ServerResponse{
		Ok: true,
		Data: HomeResponse{
			Cogs: homeResponseList,
			Chart: ChartMetadata{
				YAxisTick: yAxisTick,
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
