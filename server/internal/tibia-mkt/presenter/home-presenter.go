package presenter

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/types"
	"time"
)

type ChartMetadataResponse struct {
	XAxisTick []string     `json:"xAxisTick"`
	YAxisTick []types.Tick `json:"yAxisTick"`
}

type HomeResponse struct {
	Cogs map[string]CogSkuChartResponse `json:"cogs"`
}

type CogSkuChartResponse struct {
	Wiki         string                `json:"wiki"`
	Cog          []types.GoodResponse  `json:"cog"`
	Chart        ChartMetadataResponse `json:"chartMetadata"`
	PagePosition int8                  `json:"pagePosition"`
}

type HomePresenter struct {
	cogRepository types.Repository[types.Good]
}

func NewHomePresenter(repository types.Repository[types.Good]) *HomePresenter {
	return &HomePresenter{
		cogRepository: repository,
	}
}

func (p *HomePresenter) Format(data any) (types.ServerResponse, error) {
	cogSkuMatrix, ok := data.([][]types.GoodRecord)

	if !ok {
		return types.ServerResponse{}, types.ApiError{
			Msg:      "Assertion failed, data is not a matrix of type GoodRecord",
			Function: "Format",
			File:     "home-presenter.go",
		}
	}

	var homeResponseMap = make(map[string]CogSkuChartResponse)

	for i := 0; i < len(cogSkuMatrix); i++ {
		cogSkuList := cogSkuMatrix[i]

		var (
			buyOfferTotal      int
			sellOfferTotal     int
			cogSkuResponseList []types.GoodResponse
			highestSellPrice   types.Tick
			lowestBuyPrice     types.Tick
			yAxisDomain        []types.Tick
			xAxisDomain        []string
		)

		highestSellPrice = types.Tick{Price: cogSkuList[0].SellPrice, Date: cogSkuList[0].Date.Format(time.DateOnly)}
		lowestBuyPrice = types.Tick{Price: cogSkuList[0].BuyPrice, Date: cogSkuList[0].Date.Format(time.DateOnly)}

		for _, cogSku := range cogSkuList {
			buyOfferTotal = buyOfferTotal + cogSku.BuyPrice
			sellOfferTotal = sellOfferTotal + cogSku.SellPrice

			if highestSellPrice.Price < cogSku.SellPrice {
				highestSellPrice.Price = cogSku.SellPrice
				highestSellPrice.Date = cogSku.Date.Format(time.DateOnly)
			}

			if lowestBuyPrice.Price > cogSku.BuyPrice {
				lowestBuyPrice.Price = cogSku.BuyPrice
				lowestBuyPrice.Date = cogSku.Date.Format(time.DateOnly)
			}

			cogSkuResponseList = append(cogSkuResponseList, types.GoodResponse{
				BuyOffer:  cogSku.BuyPrice,
				SellOffer: cogSku.SellPrice,
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

		if len(cogSkuList) <= 0 {

			homeResponseMap[cogSkuList[0].ItemName] = CogSkuChartResponse{
				Wiki: cog.Link,
				Cog:  cogSkuResponseList,
				Chart: ChartMetadataResponse{
					YAxisTick: yAxisDomain,
					XAxisTick: xAxisDomain,
				},
				PagePosition: pageConfig.Position,
			}

			continue
		}

		homeResponseMap[cogSkuList[0].ItemName] = CogSkuChartResponse{
			Wiki: cog.Link,
			Cog:  cogSkuResponseList,
			Chart: ChartMetadataResponse{
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

func (p *HomePresenter) getWikiLink(itemName string) (types.Good, error) {
	var filters []types.Filter

	filters = append(filters, types.Filter{
		Name:    "name",
		Operand: constants.Equal,
		Value:   itemName,
	})

	criteria := types.Criteria{Filters: filters}

	result, err := p.cogRepository.FindOne(criteria)

	if err != nil {
		return types.Good{}, err
	}

	return result, nil
}

func (p *HomePresenter) getPagePosition(item types.Good) types.GoodConfig {
	switch item.Name {
	case constants.TibiaCoinEntity:
		return types.GoodConfig{
			CogId:    item.Id,
			Position: 1,
			Columns:  12,
			Rows:     1,
		}
	case constants.HoneycombEntity:
		return types.GoodConfig{
			CogId:    item.Id,
			Position: 2,
			Columns:  6,
			Rows:     1,
		}
	default:
		return types.GoodConfig{}
	}
}
