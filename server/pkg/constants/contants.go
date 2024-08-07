package constants

const (
	WorldSecura         = "Secura"
	HoneycombEntity     = "honeycomb"
	TibiaCoinEntity     = "tibiaCoin"
	SwamplingWoodEntity = "swamplingWood"
)

const (
	Day1  = "01"
	Day10 = "10"
	Day20 = "20"
	Day30 = "30"
	Day31 = "31"
)

const (
	SellOfferType = "SELL_OFFER"
	BuyOfferType  = "BUY_OFFER"
)

// Criteria

const (
	Equal              = "="
	GreaterThanOrEqual = ">="
	LessThanOrEqual    = "<="
)

// Errors

const (
	ServerGenericError       = "SERVER_ERROR"
	NoCogSearchParamProvided = "NO_COG_SEARCH_PARAM_PROVIDED"
)

// External API

const (
	TibiaDataApiBaseUrl           string = "https://api.tibiadata.com"
	TibiaDataApiVersion           string = "v4"
	TibiaDataApiKillStatisticsUrl string = "killstatistics"
)
