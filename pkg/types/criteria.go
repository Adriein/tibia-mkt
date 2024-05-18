package types

type Filter struct {
	Name    string
	Operand string
	Value   string
}

type Criteria struct {
	Filters []Filter
}
