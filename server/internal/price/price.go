package price

import "time"

type Price struct {
	Id         string
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
