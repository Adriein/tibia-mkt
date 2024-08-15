package types

import "time"

type GoodRecord struct {
	Id        string
	ItemName  string
	BuyPrice  int
	SellPrice int
	Date      time.Time
	World     string
}

type GoodDrop struct {
	Name     string  `json:"name"`
	DropRate float64 `json:"dropRate"`
}

type Good struct {
	Id        string
	Name      string
	Link      string
	Drop      []GoodDrop
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GoodConfig struct {
	CogId    string
	Position int8
	Columns  int8
	Rows     int8
}

type GoodRecordInterval struct {
	Name string
	From string
	To   string
}

type GoodResponse struct {
	BuyOffer  int    `json:"buyOffer"`
	SellOffer int    `json:"sellOffer"`
	Date      string `json:"date"`
	World     string `json:"world"`
}
