package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"../service"
	_ "github.com/go-sql-driver/mysql"
)

// マスタからSELECTしたデータをマッピングする構造体
type DrawResponse struct {
	Times int `json:"times"`
}

func GachaDraw(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case "POST":
		user, err := service.AuthUser(r.Header.Get("x-token"), db)
		if err != nil {
			return
		}

		var draw_response DrawResponse
		body := r.Body
		defer body.Close()
		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		// byte配列にしたbody内のjsonをgoで扱えるようにobjectに変換
		err = json.Unmarshal(buf.Bytes(), &draw_response)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(draw_response.Times)

		gacha_draw_results, err := service.GachaPlay(user, draw_response.Times, db)
		if err != nil {
			fmt.Println(err)
			return
		}

		service.RespondJSON(w, 200, gacha_draw_results)
	}
	return
}

func GetUserCharacterAll(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case "GET":
		user, err := service.AuthUser(r.Header.Get("x-token"), db)
		if err != nil {
			return
		}
		fmt.Println(user)

		service.
	}
	return
}
