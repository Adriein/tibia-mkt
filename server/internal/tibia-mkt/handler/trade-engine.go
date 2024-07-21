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
	date, _ := time.Parse(time.DateTime, "2024-12-12 00:00:00")

	result, _ := h.engine.Execute(types.CogInterval{Name: "honeycomb", From: date, To: time.Now()})

	response, _ := h.presenter.Format(result)

	if err := service.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
