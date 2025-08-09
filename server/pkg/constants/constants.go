package constants

//Env var keys

const (
	DatabaseUser     = "DATABASE_USER"
	DatabasePassword = "DATABASE_PASSWORD"
	DatabaseName     = "DATABASE_NAME"
	ServerPort       = "SERVER_PORT"
	TibiaMktApiKey   = "TIBIA_MKT_API_KEY"
	Env              = "ENV"
	Production       = "PRODUCTION"
)

// Tibia Mkt

const (
	ApiKeyHeader = "TibiaMkt-Token"
	OkResKey     = "ok"
	ErrorResKey  = "error"
	DataResKey   = "data"
)

const (
	WorldSecura               = "Secura"
	HoneycombEntity           = "honeycomb"
	TibiaCoinEntity           = "tibiaCoin"
	SwamplingWoodEntity       = "swamplingWood"
	BrokenShamanicStaffEntity = "brokenShamanicStaff"
	SellOffer                 = "SELL"
	BuyOffer                  = "BUY"
	VolatileMarketStatus      = "Volatile"
	StableMarketStatus        = "Stable"
	RiskyMarketStatus         = "Risky"
)

const (
	Day1  = "01"
	Day10 = "10"
	Day20 = "20"
	Day30 = "30"
	Day31 = "31"
)

// Criteria

const (
	Equal              = "="
	GreaterThanOrEqual = ">="
	LessThanOrEqual    = "<="
)

// Errors

const (
	ServerGenericError         = "SERVER_ERROR"
	NoGoodSearchParamProvided  = "NO_GOOD_SEARCH_PARAM_PROVIDED"
	NoWorldSearchParamProvided = "NO_WORLD_SEARCH_PARAM_PROVIDED"
)

// External API

const (
	TibiaDataApiBaseUrl           string = "https://api.tibiadata.com"
	TibiaDataApiVersion           string = "v4"
	TibiaDataApiKillStatisticsUrl string = "killstatistics"
)

const (
	IncomingTimeFormat = "20060102150405"
)

// App users

const (
	TibiaMktCronUser string = "tibia-mkt"
)
