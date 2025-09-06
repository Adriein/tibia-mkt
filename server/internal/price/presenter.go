package price

import (
	"time"

	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/gin-gonic/gin"
)

type Presenter struct {
}

type Response struct {
	Prices map[string]PricesChart `json:"prices"`
}

type PricesChart struct {
	Wiki         string     `json:"wiki"`
	BuyOffer     []PriceDto `json:"buyOffer"`
	SellOffer    []PriceDto `json:"sellOffer"`
	PagePosition int8       `json:"pagePosition"`
}

type PriceDto struct {
	UnitPrice int    `json:"unitPrice"`
	Amount    int    `json:"amount"`
	CreatedAt string `json:"createdAt"`
	World     string `json:"world"`
}

type PriceConfig struct {
	CogId    string
	Position int8
	Columns  int8
	Rows     int8
}

func NewPresenter() *Presenter {
	return &Presenter{}
}

func (p *Presenter) Format(data [][]*Price) gin.H {
	var homeResponseMap = make(map[string]PricesChart)

	for i := 0; i < len(data); i++ {
		cogSkuList := data[i]

		var (
			buyOfferResponses  []PriceDto
			sellOfferResponses []PriceDto
		)

		for _, cogSku := range cogSkuList {
			if cogSku.UnitPrice == -1 {
				continue
			}

			if cogSku.OfferType == constants.SellOffer {
				sellOfferResponses = append(sellOfferResponses, PriceDto{
					UnitPrice: cogSku.UnitPrice,
					Amount:    cogSku.GoodAmount,
					CreatedAt: cogSku.CreatedAt.Format(time.DateOnly),
					World:     cogSku.World,
				})

				continue
			}

			buyOfferResponses = append(buyOfferResponses, PriceDto{
				UnitPrice: cogSku.UnitPrice,
				Amount:    cogSku.GoodAmount,
				CreatedAt: cogSku.CreatedAt.Format(time.DateOnly),
				World:     cogSku.World,
			})
		}

		pageConfig := p.getPagePosition(cogSkuList[0])

		homeResponseMap[cogSkuList[0].GoodName] = PricesChart{
			Wiki:         "cog.Link",
			BuyOffer:     buyOfferResponses,
			SellOffer:    sellOfferResponses,
			PagePosition: pageConfig.Position,
		}
	}

	return gin.H{
		constants.OkResKey:   true,
		constants.DataResKey: homeResponseMap,
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
	case constants.TurtleShell:
		return PriceConfig{
			CogId:    item.Id,
			Position: 5,
			Columns:  6,
			Rows:     1,
		}
	case constants.CobraRod:
		return PriceConfig{
			CogId:    item.Id,
			Position: 6,
			Columns:  6,
			Rows:     1,
		}
	default:
		return PriceConfig{}
	}
}
