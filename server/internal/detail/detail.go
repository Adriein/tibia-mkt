package detail

type DetailInsights struct {
	MarketType   string `json:"marketType"`
	BuyPressure  int    `json:"buyPressure"`
	SellPressure int    `json:"sellPressure"`
	Liquidity    int    `json:"liquidity"`
}
type DetailOverview struct {
	BuySellSpread                  int    `json:"buySellSpread"`
	SpreadPercentage               int    `json:"spreadPercentage"`
	MarketCap                      int    `json:"marketCap"`
	LastTwentyFourHoursVolume      int    `json:"lastTwentyFourHoursVolume"`
	MarketStatus                   string `json:"marketStatus"`
	MarketVolumePercentageTendency int    `json:"marketVolumePercentageTendency"`
	TotalGoodsBeingSold            int    `json:"totalGoodsBeingSold"`
}
type DetailStats struct {
	SellOffersMean         int `json:"sellOffersMean"`
	SellOffersStdDeviation int `json:"sellOffersStdDeviation"`
	SellOffersMedian       int `json:"sellOffersMedian"`
	BuyOffersMean          int `json:"buyOffersMean"`
	BuyOffersStdDeviation  int `json:"buyOffersStdDeviation"`
	BuyOffersMedian        int `json:"buyOffersMedian"`
}
type Detail struct {
	Stats    DetailStats    `json:"stats"`
	Overview DetailOverview `json:"overview"`
	Insights DetailInsights `json:"insights"`
}
