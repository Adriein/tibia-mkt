package price

import (
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
		items := ctx.QueryArray("item")

		if len(items) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "item to query is mandatory"})
			return
		}

		var prices []*Price

		for _, item := range items {
			price, err := c.service.GetPrice(item)

			if err != nil {
				ctx.Error(err)
				return
			}

			prices = append(prices, price)
		}

		response := c.presenter.Format(prices)

		ctx.JSON(http.StatusOK, response)
	}
}
