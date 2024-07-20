package trade_engine

import "github.com/adriein/tibia-mkt/pkg/types"

type TradeEngineConfig struct {
	OfferCreationFeePercentage int //2%
	OfferCreationMaxFee        int //250000
	OfferCreationMinFee        int //20
	MaxSimultaneousOperations  int //100
	OfferDurationInDays        int //30
	Algorithm                  types.TradeEngineAlgorithm
}

func NewConfig() *TradeEngineConfig {
	return &TradeEngineConfig{
		OfferCreationFeePercentage: 2,
		OfferCreationMaxFee:        250_000,
		OfferCreationMinFee:        20,
		MaxSimultaneousOperations:  100,
		OfferDurationInDays:        30,
	}
}

func (config *TradeEngineConfig) WithAlgorithm(algorithm types.TradeEngineAlgorithm) *TradeEngineConfig {
	config.Algorithm = algorithm

	return config
}
