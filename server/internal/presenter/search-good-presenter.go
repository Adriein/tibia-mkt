package presenter

import (
	"github.com/adriein/tibia-mkt/pkg/types"
)

type SearchGoodPresenter struct{}

func NewSearchGoodPresenter() *SearchGoodPresenter {
	return &SearchGoodPresenter{}
}

func (p *SearchGoodPresenter) Format(data any) (types.ServerResponse, error) {
	input, ok := data.([]types.Good)

	if !ok {
		return types.ServerResponse{}, types.ApiError{
			Msg:      "Assertion failed, data is not an array of Good",
			Function: "Format",
			File:     "search-good-presenter.go",
		}
	}

	result := make([]string, len(input))

	for index, good := range input {
		result[index] = good.Name
	}

	response := types.ServerResponse{
		Ok:   true,
		Data: result,
	}

	return response, nil
}
