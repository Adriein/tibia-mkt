package types

type Presenter interface {
	Format(data any) ([]byte, error)
}
