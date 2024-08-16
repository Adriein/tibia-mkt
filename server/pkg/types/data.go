package types

type DataSnapshot struct {
	Id           string  `json:"id"`
	Cog          string  `json:"cog"`
	StdDeviation float64 `json:"stdDeviation"`
	Mean         int     `json:"mean"`
	TotalDropped int     `json:"totalDropped"`
	ExecutedBy   string  `json:"executedBy"`
	CreatedAt    string  `json:"createdAt"`
	UpdatedAt    string  `json:"updatedAt"`
}

type SeedRequest struct {
	Items []SeedGood `json:"items"`
}

type SeedGood struct {
	Name      string         `json:"name"`
	Wiki      string         `json:"wiki"`
	Creatures []SeedCreature `json:"creatures"`
}

type SeedCreature struct {
	Name     string  `json:"name"`
	DropRate float64 `json:"dropRate"`
}
