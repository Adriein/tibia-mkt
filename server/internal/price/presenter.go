package price

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/gin-gonic/gin"
	"time"
)

type Presenter struct {
}

type ChartMetadataDto struct {
	XAxisTick []string    `json:"xAxisTick"`
	YAxisTick []ChartTick `json:"yAxisTick"`
}

type Response struct {
	Prices map[string]PricesChart `json:"prices"`
}

type PricesChart struct {
	Wiki         string           `json:"wiki"`
	Prices       []PriceDto       `json:"prices"`
	Chart        ChartMetadataDto `json:"chartMetadata"`
	PagePosition int8             `json:"pagePosition"`
}

type PriceDto struct {
	BuyOffer  int    `json:"buyOffer"`
	SellOffer int    `json:"sellOffer"`
	Amount    int    `json:"amount"`
	Date      string `json:"date"`
	World     string `json:"world"`
}

type PriceConfig struct {
	CogId    string
	Position int8
	Columns  int8
	Rows     int8
}

type ChartTick struct {
	Price int    `json:"price"`
	Date  string `json:"date"`
}

func NewPresenter() *Presenter {
	return &Presenter{}
}

func (p *Presenter) Format(data [][]*Price) gin.H {
	var homeResponseMap = make(map[string]PricesChart)

	for i := 0; i < len(data); i++ {
		cogSkuList := data[i]

		var (
			buyOfferTotal      int
			sellOfferTotal     int
			cogSkuResponseList []PriceDto
			highestSellPrice   ChartTick
			lowestBuyPrice     ChartTick
			yAxisDomain        []ChartTick
			xAxisDomain        []string
		)

		highestSellPrice = ChartTick{Price: cogSkuList[0].SellPrice, Date: cogSkuList[0].RegisteredAt.Format(time.DateOnly)}
		lowestBuyPrice = ChartTick{Price: cogSkuList[0].BuyPrice, Date: cogSkuList[0].RegisteredAt.Format(time.DateOnly)}

		for _, cogSku := range cogSkuList {
			buyOfferTotal = buyOfferTotal + cogSku.BuyPrice
			sellOfferTotal = sellOfferTotal + cogSku.SellPrice

			if highestSellPrice.Price < cogSku.SellPrice {
				highestSellPrice.Price = cogSku.SellPrice
				highestSellPrice.Date = cogSku.RegisteredAt.Format(time.DateOnly)
			}

			if lowestBuyPrice.Price > cogSku.BuyPrice {
				lowestBuyPrice.Price = cogSku.BuyPrice
				lowestBuyPrice.Date = cogSku.RegisteredAt.Format(time.DateOnly)
			}

			cogSkuResponseList = append(cogSkuResponseList, PriceDto{
				BuyOffer:  cogSku.BuyPrice,
				SellOffer: cogSku.SellPrice,
				Amount:    0,
				Date:      cogSku.RegisteredAt.Format(time.DateOnly),
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

		pageConfig := p.getPagePosition(cogSkuList[0])

		if len(cogSkuList) <= 0 {
			homeResponseMap[cogSkuList[0].GoodName] = PricesChart{
				Wiki:   "cog.Link",
				Prices: cogSkuResponseList,
				Chart: ChartMetadataDto{
					YAxisTick: yAxisDomain,
					XAxisTick: xAxisDomain,
				},
				PagePosition: pageConfig.Position,
			}

			continue
		}

		homeResponseMap[cogSkuList[0].GoodName] = PricesChart{
			Wiki:   "cog.Link",
			Prices: cogSkuResponseList,
			Chart: ChartMetadataDto{
				YAxisTick: yAxisDomain,
				XAxisTick: xAxisDomain,
			},
			PagePosition: pageConfig.Position,
		}
	}

	return gin.H{
		"ok":   true,
		"data": homeResponseMap,
	}
}

func (p *Presenter) getPagePosition(item *Price) PriceConfig {
	switch item.GoodName {
	case constants.TibiaCoinEntity:
		return PriceConfig{
			CogId:    item.Id,
			Position: 1,
			Columns:  12,
			Rows:     1,
		}
	case constants.HoneycombEntity:
		return PriceConfig{
			CogId:    item.Id,
			Position: 2,
			Columns:  6,
			Rows:     1,
		}
	case constants.SwamplingWoodEntity:
		return PriceConfig{
			CogId:    item.Id,
			Position: 3,
			Columns:  6,
			Rows:     1,
		}
	case constants.BrokenShamanicStaffEntity:
		return PriceConfig{
			CogId:    item.Id,
			Position: 4,
			Columns:  6,
			Rows:     1,
		}
	default:
		return PriceConfig{}
	}
}
