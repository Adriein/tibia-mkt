package event

import "time"

type Event struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	GoodName    string    `json:"goodName"`
	World       string    `json:"world"`
	Description string    `json:"description"`
	OccurredAt  time.Time `json:"occurredAt"`
}
