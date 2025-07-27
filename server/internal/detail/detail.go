package detail

type Detail struct {
	SellOffersMean         int `json:"sellOffersMean"`
	SellOffersStdDeviation int `json:"sellOffersStdDeviation"`
	SellOffersMedian       int `json:"sellOffersMedian"`
	BuyOffersMean          int `json:"buyOffersMean"`
	BuyOffersStdDeviation  int `json:"buyOffersStdDeviation"`
	BuyOffersMedian        int `json:"buyOffersMedian"`
}
