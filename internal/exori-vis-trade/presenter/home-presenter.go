package presenter

import (
	"encoding/json"
	"github.com/adriein/exori-vis-trade/pkg/types"
)

type HomePresenter struct {
	data any
}

func NewHomePresenter(data any) *HomePresenter {
	return &HomePresenter{
		data: data,
	}
}

func (p *HomePresenter) Format() ([]byte, error) {
	cogSkuList, ok := p.data.([]types.CogSku)

	if !ok {
		return nil, types.ApiError{
			Msg:      "Assertion Failed",
			Function: "Format",
			File:     "home-presenter.go",
		}
	}

	for _, cogSku := range cogSkuList {

	}

	response, jsonErr := json.Marshal(p.data)

	if jsonErr != nil {
		return nil, types.ApiError{
			Msg:      jsonErr.Error(),
			Function: "Format",
			File:     "home-presenter.go",
		}
	}

	return response, nil
}
