package detail

type DetailOverview struct {
	BuySellSpread                  int    `json:"buySellSpread"`
	SpreadPercentage               int    `json:"spreadPercentage"`
	MarketCap                      int    `json:"marketCap"`
	LastTwentyFourHoursVolume      int    `json:"lastTwentyFourHoursVolume"`
	MarketStatus                   string `json:"marketStatus"`
	MarketVolumeTendency           string `json:"marketVolumeTendency"`
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
}
