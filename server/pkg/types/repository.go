package types

type GoodRecordRepository interface {
	GoodName() string
	Find(criteria Criteria) ([]GoodRecord, error)
	Save(entity GoodRecord) error
}

type Repository[T any] interface {
	Find(criteria Criteria) ([]T, error)
	FindOne(criteria Criteria) (T, error)
	Save(entity T) error
}
