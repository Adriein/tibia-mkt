package types

type CogRepository interface {
	Find(criteria Criteria) ([]CogSku, error)
	Save(entity CogSku) error
}
