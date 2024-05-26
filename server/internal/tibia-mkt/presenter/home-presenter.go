package presenter

import (
	"encoding/json"
	"github.com/adriein/tibia-mkt/pkg/types"
	"time"
)

type HomeResponseCogSku struct {
	BuyPrice  int    `json:"buyPrice"`
	SellPrice int    `json:"sellPrice"`
	Date      string `json:"date"`
	World     string `json:"world"`
}

type HomePresenter struct{}

func NewHomePresenter() *HomePresenter {
	return &HomePresenter{}
}

func (p *HomePresenter) Format(data any) ([]byte, error) {
	cogSkuList, ok := data.([]types.CogSku)
	var homeResponseList []HomeResponseCogSku

	if !ok {
		return nil, types.ApiError{
			Msg:      "Assertion failed, data is not of type CogSku",
			Function: "Format",
			File:     "home-presenter.go",
		}
	}

	for _, cogSku := range cogSkuList {
		homeResponseList = append(homeResponseList, HomeResponseCogSku{
			BuyPrice:  cogSku.BuyPrice,
			SellPrice: cogSku.SellPrice,
			Date:      cogSku.Date.Format(time.DateOnly),
			World:     cogSku.World,
		})
	}

	response := &types.ServerResponse{
		Ok:   true,
		Data: homeResponseList,
	}

	bytes, jsonErr := json.Marshal(response)

	if jsonErr != nil {
		return nil, types.ApiError{
			Msg:      jsonErr.Error(),
			Function: "Format",
			File:     "home-presenter.go",
		}
	}

	return bytes, nil
}
