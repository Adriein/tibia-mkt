package presenter

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/types"
	"time"
)

type CogSkuResponse struct {
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
	Cogs map[string]CogSkuChartResponse `json:"cogs"`
}

type CogSkuChartResponse struct {
	Wiki  string           `json:"wiki"`
	Cog   []CogSkuResponse `json:"cog"`
	Chart ChartMetadata    `json:"chartMetadata"`
}

type HomePresenter struct {
	cogRepository types.Repository
}

func NewHomePresenter(repository types.Repository) *HomePresenter {
	return &HomePresenter{
		cogRepository: repository,
	}
}

func (p *HomePresenter) Format(data any) (types.ServerResponse, error) {
	cogSkuMatrix, ok := data.([][]types.CogSku)

	if !ok {
		return types.ServerResponse{}, types.ApiError{
			Msg:      "Assertion failed, data is not a matrix of type CogSku",
			Function: "Format",
			File:     "home-presenter.go",
		}
	}

	var homeResponseMap = make(map[string]CogSkuChartResponse)

	for i := 0; i < len(cogSkuMatrix); i++ {
		cogSkuList := cogSkuMatrix[i]

		var (
			cogSkuResponseList []CogSkuResponse
			highestSellPrice   int
			lowestBuyPrice     int
			yAxisDomain        []int
			xAxisDomain        []string
		)

		highestSellPrice = cogSkuList[0].SellPrice
		lowestBuyPrice = cogSkuList[0].BuyPrice

		for _, cogSku := range cogSkuList {
			if highestSellPrice < cogSku.SellPrice {
				highestSellPrice = cogSku.SellPrice
			}

			if lowestBuyPrice > cogSku.SellPrice {
				lowestBuyPrice = cogSku.BuyPrice
			}

			cogSkuResponseList = append(cogSkuResponseList, CogSkuResponse{
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

		cog, err := p.getWikiLink(cogSkuList[0].ItemName)

		if err != nil {
			return types.ServerResponse{}, err
		}

		homeResponseMap[cogSkuList[0].ItemName] = CogSkuChartResponse{
			Wiki: cog.Link,
			Cog:  cogSkuResponseList,
			Chart: ChartMetadata{
				YAxisTick: yAxisDomain,
				XAxisTick: xAxisDomain,
			},
		}
	}

	response := types.ServerResponse{
		Ok:   true,
		Data: homeResponseMap,
	}

	return response, nil
}

func (p *HomePresenter) getWikiLink(itemName string) (types.Cog, error) {
	var filters []types.Filter

	filters = append(filters, types.Filter{
		Name:    "name",
		Operand: "=",
		Value:   itemName,
	})

	criteria := types.Criteria{Filters: filters}

	result, err := p.cogRepository.FindOne(criteria)

	if err != nil {
		return types.Cog{}, err
	}

	return result, nil
}
