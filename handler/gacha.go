package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"../service"
	_ "github.com/go-sql-driver/mysql"
)

type DrawResponse struct {
	Times int `json:"times"`
}

func (handler *UserHandler) GachaDraw(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		user, err := service.AuthUser(r.Header.Get("x-token"), handler.DB)
		if err != nil {
			return
		}

		var draw_response DrawResponse
		body := r.Body
		defer body.Close()
		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		err = json.Unmarshal(buf.Bytes(), &draw_response)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(draw_response.Times)

		gacha_draw_results, err := service.GachaPlay(user, draw_response.Times, handler.DB)
		if err != nil {
			fmt.Println(err)
			return
		}

		service.RespondJSON(w, 200, gacha_draw_results)
	}
	return
}

func (handler *UserHandler) GetUserCharacterAll(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		user, err := service.AuthUser(r.Header.Get("x-token"), handler.DB)
		if err != nil {
			return
		}
		list_user_character, err := service.GetUserCharacters(user.ID, handler.DB)
		if err != nil {
			fmt.Println(err)
			return
		}

		service.RespondJSON(w, 200, list_user_character)
	}
	return
}
