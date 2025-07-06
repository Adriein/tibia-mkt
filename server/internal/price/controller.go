package price

import (
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

func (c *Controller) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		item := ctx.Query("good")

		if item == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "item to query is mandatory"})
			return
		}

		c.service.GetPrice(item)

		ctx.JSON(http.StatusOK, gin.H{"ok": true, "data": "pong"})
	}
}
