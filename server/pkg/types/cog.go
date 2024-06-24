package types

import "time"

type CogSku struct {
	Id        string
	ItemName  string
	BuyPrice  int
	SellPrice int
	Date      time.Time
	World     string
}
