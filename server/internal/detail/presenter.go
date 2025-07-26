package detail

import (
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/gin-gonic/gin"
)

type Presenter struct {
}

func NewPresenter() *Presenter {
	return &Presenter{}
}

func (p *Presenter) Format(detail *Detail) gin.H {
	return gin.H{
		constants.OkResKey:   true,
		constants.DataResKey: detail,
	}
}
