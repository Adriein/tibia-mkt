package handler

import (
	"github.com/adriein/tibia-mkt/internal/trade-engine"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
)

type TradeEngineRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
	Item string `json:"item"`
}

type TradeEngineHandler[T any] struct {
	engine    *trade_engine.TradeEngine[T]
	presenter types.Presenter
}

func NewTradeEngineHandler[T any](
	engine *trade_engine.TradeEngine[T],
	presenter types.Presenter,
) *TradeEngineHandler[T] {
	return &TradeEngineHandler[T]{
		engine:    engine,
		presenter: presenter,
	}
}

func (h *TradeEngineHandler[T]) Handler(w http.ResponseWriter, r *http.Request) error {
	tradeEngineRequest, decodeErr := service.Decode[TradeEngineRequest](r)

	if decodeErr != nil {
		response := types.ServerResponse{
			Ok: false,
		}

		if err := service.Encode[types.ServerResponse](w, http.StatusInternalServerError, response); err != nil {
			return err
		}

		return types.ApiError{
			Msg:      decodeErr.Error(),
			Function: "Handler",
			File:     "trade-engine.go",
		}
	}

	result, tradeEngineErr := h.engine.Execute(
		types.CogInterval{
			Name: tradeEngineRequest.Item,
			From: tradeEngineRequest.From,
			To:   tradeEngineRequest.To,
		},
	)

	if tradeEngineErr != nil {
		response := types.ServerResponse{
			Ok: false,
		}

		if err := service.Encode[types.ServerResponse](w, http.StatusInternalServerError, response); err != nil {
			return err
		}

		return tradeEngineErr
	}

	response, presenterErr := h.presenter.Format(result)

	if presenterErr != nil {
		response := types.ServerResponse{
			Ok:    false,
			Error: constants.ServerGenericError,
		}

		if err := service.Encode[types.ServerResponse](w, http.StatusInternalServerError, response); err != nil {
			return err
		}

		return presenterErr
	}

	if err := service.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
