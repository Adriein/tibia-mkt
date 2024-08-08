package types

type CogRepository interface {
	EntityName() string
	Find(criteria Criteria) ([]CogSku, error)
	Save(entity CogSku) error
}

type Repository[T any] interface {
	Find(criteria Criteria) ([]T, error)
	FindOne(criteria Criteria) (T, error)
	Save(entity T) error
}
