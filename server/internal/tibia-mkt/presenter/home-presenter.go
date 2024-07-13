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

type YAxisTick struct {
	Price int    `json:"price"`
	Date  string `json:"date"`
}

type ChartMetadata struct {
	XAxisTick []string    `json:"xAxisTick"`
	YAxisTick []YAxisTick `json:"yAxisTick"`
}

type HomeResponse struct {
	Cogs map[string]CogSkuChartResponse `json:"cogs"`
}

type CogSkuChartResponse struct {
	Wiki         string           `json:"wiki"`
	Cog          []CogSkuResponse `json:"cog"`
	Chart        ChartMetadata    `json:"chartMetadata"`
	PagePosition int8             `json:"pagePosition"`
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
			highestSellPrice   YAxisTick
			lowestBuyPrice     YAxisTick
			yAxisDomain        []YAxisTick
			xAxisDomain        []string
		)

		highestSellPrice = YAxisTick{Price: cogSkuList[0].SellPrice, Date: cogSkuList[0].Date.Format(time.DateOnly)}
		lowestBuyPrice = YAxisTick{Price: cogSkuList[0].BuyPrice, Date: cogSkuList[0].Date.Format(time.DateOnly)}

		for _, cogSku := range cogSkuList {
			if highestSellPrice.Price < cogSku.SellPrice {
				highestSellPrice.Price = cogSku.SellPrice
				highestSellPrice.Date = cogSku.Date.Format(time.DateOnly)
			}

			if lowestBuyPrice.Price > cogSku.BuyPrice {
				lowestBuyPrice.Price = cogSku.BuyPrice
				lowestBuyPrice.Date = cogSku.Date.Format(time.DateOnly)
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

		pageConfig := p.getPagePosition(cog)

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
			PagePosition: pageConfig.Position,
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

func (p *HomePresenter) getPagePosition(item types.Cog) types.CogConfig {
	switch item.Name {
	case constants.TibiaCoinEntity:
		return types.CogConfig{
			CogId:    item.Id,
			Position: 1,
			Columns:  12,
			Rows:     1,
		}
	case constants.HoneycombEntity:
		return types.CogConfig{
			CogId:    item.Id,
			Position: 2,
			Columns:  6,
			Rows:     1,
		}
	default:
		return types.CogConfig{}
	}
}
