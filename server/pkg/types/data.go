package types

type DataSnapshot struct {
	Id           string  `json:"id"`
	Cog          string  `json:"cog"`
	StdDeviation float64 `json:"stdDeviation"`
	Mean         int     `json:"mean"`
	ExecutedBy   string  `json:"executedBy"`
	CreatedAt    string  `json:"createdAt"`
	UpdatedAt    string  `json:"updatedAt"`
}
