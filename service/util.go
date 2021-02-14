package service

import (

	"net/http"
	"../schema"
	"encoding/json"
	"io"
	"bytes"


	"github.com/google/uuid"
)

func GainUserName(r *http.Request) (name string, err error) {
	var user schema.User
	body := r.Body
	defer body.Close()
	buf := new(bytes.Buffer)
	io.Copy(buf, body)

	if err := json.Unmarshal(buf.Bytes(), &user); err != nil {
		return "", err
	}
	return user.Name, nil
}


func CreateUuid() (token string, err error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	uu := u.String()
	return uu, err
}
