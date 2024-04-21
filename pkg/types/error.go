package types

import "fmt"

type EvtError struct {
	Msg      string
	Function string
	File     string
}

func (e *EvtError) Error() string {
	return fmt.Sprintf("Error %s in %s at %s", e.Msg, e.File, e.Function)
}
