package script

import "time"

type RawJsonObj struct {
	Id         int     `json:"id"`
	Time       float64 `json:"time"`
	BuyOffer   int     `json:"buy_offer"`
	SellOffer  int     `json:"sell_offer"`
	BuyOffers  int     `json:"buy_offers"`
	SellOffers int     `json:"sell_offers"`
}

type JsonObj struct {
	Time       time.Time
	BuyOffer   int
	SellOffer  int
	BuyOffers  int
	SellOffers int
}
