package types

import "time"

type TradeEngineOperation struct {
	Id        string
	OfferType string
	OfferDate time.Time
	Price     int
	MarketFee int
}

type TradeEngineResult struct {
	IntervalPriceAverage int
	Roi                  int //%
	WinRate              int //%
	SharpeRatio          int //%
	Invested             int
	Profit               int
	Operations           []TradeEngineOperation
}

type TradeEngineAlgorithm interface {
	Apply([]CogSku) error
}
