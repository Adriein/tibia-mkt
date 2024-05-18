package types

import "time"

type CogSku struct {
	Id     string
	Price  int
	Date   time.Time
	World  string
	Action string
}
