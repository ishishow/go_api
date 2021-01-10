package handler

import (
	"database/sql"
	"net/http"

	"../model"
	"../service"
	_ "github.com/go-sql-driver/mysql"
)

// マスタからSELECTしたデータをマッピングする構造体

func GachaDraw(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case "POST":
		var user model.User

		user, err := service.AuthUser(r.Header.Get("x-token"), db)
		if err != nil {
			return
		}

		character_ids, err := service.GachaPlay(user, r.Header.Get("times"), db)
		if err != nil {
			return
		}

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
