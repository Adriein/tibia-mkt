package types

type CogRepository interface {
	EntityName() string
	Find(criteria Criteria) ([]CogSku, error)
	Save(entity CogSku) error
}
