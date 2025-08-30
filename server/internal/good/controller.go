package good

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
		item := ctx.Query("item")

		if item == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "item to query is mandatory"})
			return
		}

		good, err := c.service.GetGood(item)

		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"ok": true, "data": good})
	}
}
