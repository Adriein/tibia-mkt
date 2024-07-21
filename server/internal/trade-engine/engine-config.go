package trade_engine

type TradeEngineConfig struct {
	OfferCreationFeePercentage int //2%
	OfferCreationMaxFee        int //250000
	OfferCreationMinFee        int //20
	MaxSimultaneousOperations  int //100
	OfferDurationInDays        int //30
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
