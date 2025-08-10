package price

import "time"

type Price struct {
	Id         string
	BatchId    int
	MarketId   string
	OfferType  string
	GoodName   string
	World      string
	CreatedBy  string
	GoodAmount int
	UnitPrice  int
	TotalPrice int
	EndAt      time.Time
	CreatedAt  time.Time
}
