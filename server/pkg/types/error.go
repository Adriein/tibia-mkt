package types

import "fmt"

type ApiErrorInterface interface {
	Error() string
	IsDomain() bool
	PresentableError() string
}

type ApiError struct {
	Msg      string
	Function string
	File     string
	Domain   bool
}

func (e ApiError) Error() string {
	return fmt.Sprintf("Error %s in file %s at function %s", e.Msg, e.File, e.Function)
}

func (e ApiError) IsDomain() bool {
	return e.Domain
}

func (e ApiError) PresentableError() string {
	return fmt.Sprintf("%s", e.Msg)
}
