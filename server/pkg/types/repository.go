package types

type CogRepository interface {
	EntityName() string
	Find(criteria Criteria) ([]CogSku, error)
	Save(entity CogSku) error
}

type Repository interface {
	FindOne(criteria Criteria) (Cog, error)
	Save(entity Cog) error
}
