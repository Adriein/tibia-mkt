package types

type Detail struct {
	Wiki                  string
	Cog                   []GoodRecord
	Creatures             []CreatureKillStatistic
	SellPriceMean         int
	StdDeviation          float64
	SellOfferFrequency    []SellOfferFrequency
	SellOfferHistoricData []DataSnapshot
}

type SellOfferFrequency struct {
	Range       string  `json:"range"`
	Occurrences int     `json:"occurrences"`
	Frequency   float64 `json:"frequency"`
}
