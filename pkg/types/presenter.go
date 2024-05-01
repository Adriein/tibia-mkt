package types

type Presenter interface {
	Format() ([]byte, error)
}
