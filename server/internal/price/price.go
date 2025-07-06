package price

import "time"

type Price struct {
	Id           string
	GoodName     string
	World        string
	BuyPrice     int
	SellPrice    int
	RegisteredAt time.Time
}
