package detail

type Detail struct {
	SellOffersMean         int     `json:"sellOffersMean"`
	SellOffersStdDeviation float64 `json:"sellOffersStdDeviation"`
	SellOffersMedian       int     `json:"sellOffersMedian"`
	BuyOffersMean          int     `json:"buyOffersMean"`
	BuyOffersStdDeviation  float64 `json:"buyOffersStdDeviation"`
	BuyOffersMedian        int     `json:"buyOffersMedian"`
}
