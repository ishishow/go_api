package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"../service"
	"../db"
	"../usecase"
	_ "github.com/go-sql-driver/mysql"
)

type DrawRequest struct {
	Times int `json:"times"`
}

func (handler *UserHandler) GachaDraw(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handler.Mysql)
	user, err := usecase.Get(ctx, r.Header.Get("x-token"))
	if err != nil {
		service.RespondJSON(w, 500, err)
		return
	}


	var draw_request DrawRequest
	body := r.Body
	defer body.Close()
	buf := new(bytes.Buffer)
	io.Copy(buf, body)
	if err := json.Unmarshal(buf.Bytes(), &draw_request); err != nil {
		service.RespondJSON(w, 500, err)
		return
	}
	fmt.Println(draw_request.Times)


	gacha_draw_results, err := usecase.PlayGacha(ctx, user.ID, draw_request.Times)
	if err != nil {
		fmt.Println(err)
		return
	}

	service.RespondJSON(w, 200, gacha_draw_results)

	return
}

func (handler *UserHandler) GetUserCharacterAll(w http.ResponseWriter, r *http.Request) {
	user, err := service.AuthUser(r.Header.Get("x-token"), handler.Mysql.DB)
	if err != nil {
		return
	}
	list_user_character, err := service.GetUserCharacters(user.ID, handler.Mysql.DB)
	if err != nil {
		fmt.Println(err)
		return
	}

	service.RespondJSON(w, 200, list_user_character)
	return
}
