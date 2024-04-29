package types

type CogRepository interface {
	Find(criteria Criteria) []CogSku
}
