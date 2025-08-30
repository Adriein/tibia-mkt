package event

import (
	"net/http"

	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/gin-gonic/gin"
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
		good := ctx.Query("good")
		world := ctx.Query("world")

		if world == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				constants.OkResKey:    false,
				constants.ErrorResKey: constants.NoWorldSearchParamProvided,
			})

			return
		}

		if good == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				constants.OkResKey:    false,
				constants.ErrorResKey: constants.NoGoodSearchParamProvided,
			})

			return
		}

		events, err := c.service.GetEvents(world, good)

		if err != nil {
			ctx.Error(err)

			return
		}

		response := c.presenter.Format(events)

		ctx.JSON(http.StatusOK, response)
	}
}
