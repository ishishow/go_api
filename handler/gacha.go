package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"../model"
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
		var user model.User
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

		gacha_draw_result, err := service.GachaPlay(user, draw_response.Times, db)
		if err != nil {
			fmt.Println(err)
			return
		}
		service.RespondJSON(w, 200, gacha_draw_result)
	}
	return
}

func GetUserCharacterAll(w http.ResponseWriter, r *http.Request, db *sql.DB) (id int64, err error) {
	switch r.Method {
	case "POST":
		var user model.User

		user.Name, err = service.GainUserName(r)
		if err != nil {
			return 0, err
		}

		user.Token, err = service.CreateUuid()
		if err != nil {
			return 0, err
		}

		stmt, err := db.Prepare("INSERT INTO users(name, token, created_at, updated_at) VALUES(?, ?, now(), now())")
		if err != nil {
			return 0, err
		}
		defer stmt.Close()

		//クエリ実行
		_, err = stmt.Exec(user.Name, user.Token)
		if err != nil {
			return 0, err
		}
	}
	return 0, nil
}
