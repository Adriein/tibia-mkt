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

type CogConfig struct {
	CogId    string
	Position int8
	Columns  int8
	Rows     int8
}

type CogInterval struct {
	Name string
	From string
	To   string
}
