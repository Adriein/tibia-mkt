package script

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) SeedPrices() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if err := c.service.SeedPricesFromExternalApiJson(); err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{constants.OkResKey: true})
	}
}
