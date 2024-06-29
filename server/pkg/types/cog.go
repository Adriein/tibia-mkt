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

type Cog struct {
	Id        string
	Name      string
	Link      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
