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

type CogCreature struct {
	Name     string  `json:"name"`
	DropRate float64 `json:"dropRate"`
}

type Cog struct {
	Id        string
	Name      string
	Link      string
	Creatures []CogCreature
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

type CogSkuResponse struct {
	BuyOffer  int    `json:"buyOffer"`
	SellOffer int    `json:"sellOffer"`
	Date      string `json:"date"`
	World     string `json:"world"`
}
