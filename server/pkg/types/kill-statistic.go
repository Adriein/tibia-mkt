package types

type KillStatistic struct {
	Id           string  `json:"id"`
	CreatureName string  `json:"creatureName"`
	AmountKilled int     `json:"amountKilled"`
	DropRate     float64 `json:"dropRate"`
	ExecutedBy   string  `json:"executedBy"`
	CreatedAt    string  `json:"createdAt"`
	UpdatedAt    string  `json:"updatedAt"`
}
