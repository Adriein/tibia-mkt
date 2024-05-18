package types

import (
	"os/exec"
)

type Uuid struct {
	value []byte
}

func NewUuid() (*Uuid, error) {
	uuid, err := exec.Command("uuidgen").Output()

	if err != nil {
		return &Uuid{}, ApiError{
			Msg:      err.Error(),
			Function: "Uuid -> exec.Command().Output()",
			File:     "uuid.go",
		}
	}

	return &Uuid{value: uuid}, nil
}

func (u *Uuid) String() string {
	return string(u.value)
}
