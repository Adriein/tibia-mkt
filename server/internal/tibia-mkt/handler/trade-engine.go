package handler

import (
	"github.com/adriein/tibia-mkt/internal/trade-engine"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
	"time"
)

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

func (h *TradeEngineHandler[T]) Handler(w http.ResponseWriter, _ *http.Request) error {
	date1, _ := time.Parse(time.DateOnly, "2023-11-07")
	date2, _ := time.Parse(time.DateOnly, "2024-07-22")

	result, _ := h.engine.Execute(types.CogInterval{Name: "honeycomb", From: date1, To: date2})

	response, _ := h.presenter.Format(result)

	if err := service.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
