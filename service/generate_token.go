package service

import (
	"fmt"

	"github.com/google/uuid"
)

func CreateUuid() (token string, err error) {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return
	}
	uu := u.String()
	return uu, err
}
