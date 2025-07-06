package good

import (
	"github.com/rotisserie/eris"
	"time"
)

var (
	GoodNotFoundError = eris.New("Good not found")
)

type Creature struct {
	Name     string  `json:"name"`
	DropRate float64 `json:"dropRate"`
}

type Good struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	WikiLink  string     `json:"wikiLink"`
	Creatures []Creature `json:"creatures"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}
