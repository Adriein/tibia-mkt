package constants

//Env var keys

const (
	DatabaseDsn    = "DATABASE_URL"
	ServerPort     = "SERVER_PORT"
	TibiaMktApiKey = "TIBIA_MKT_API_KEY"
)

// Tibia Mkt

const (
	ApiKeyHeader = "TibiaMkt-Token"
)

const (
	WorldSecura               = "Secura"
	HoneycombEntity           = "honeycomb"
	TibiaCoinEntity           = "tibiaCoin"
	SwamplingWoodEntity       = "swamplingWood"
	BrokenShamanicStaffEntity = "brokenShamanicStaff"
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
	ServerGenericError        = "SERVER_ERROR"
	NoGoodSearchParamProvided = "NO_GOOD_SEARCH_PARAM_PROVIDED"
)

// External API

const (
	TibiaDataApiBaseUrl           string = "https://api.tibiadata.com"
	TibiaDataApiVersion           string = "v4"
	TibiaDataApiKillStatisticsUrl string = "killstatistics"
)

// App users

const (
	TibiaMktCronUser string = "tibia-mkt"
)
