package price

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	service   *Service
	presenter *Presenter
}

func NewController(service *Service, presenter *Presenter) *Controller {
	return &Controller{
		service:   service,
		presenter: presenter,
	}
}

func (c *Controller) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		goods := ctx.QueryArray("good")
		world := ctx.Query("world")

		if len(goods) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				constants.OkResKey:    false,
				constants.ErrorResKey: constants.NoGoodSearchParamProvided,
			})

			return
		}

		if world == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				constants.OkResKey:    false,
				constants.ErrorResKey: constants.NoWorldSearchParamProvided,
			})

			return
		}

		var prices [][]*Price

		for _, good := range goods {
			price, err := c.service.GetPrices(world, good)

			if err != nil {
				_ = ctx.Error(err)
				return
			}

			if price == nil {
				continue
			}

			prices = append(prices, price)
		}

		response := c.presenter.Format(prices)

		ctx.JSON(http.StatusOK, response)
	}
}
