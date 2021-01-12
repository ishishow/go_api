package handler // 独自のクエリパッケージ

import (
	"database/sql"
	"fmt"
	"net/http"

	"../model"
	"../service"
	_ "github.com/go-sql-driver/mysql"
)

// マスタからSELECTしたデータをマッピングする構造体

func CreateUser(w http.ResponseWriter, r *http.Request, db *sql.DB) (id int64, err error) {
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

func GetUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case "GET":
		user, err := service.AuthUser(r.Header.Get("x-token"), db)
		if err != nil {
			return
		}
		fmt.Println(user.Name)
		service.RespondJSON(w, 200, user)
	}
	return
}

func UpdateUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// プリペアードステートメント
	switch r.Method {
	case "PUT":

		name, err := service.GainUserName(r)
		if err != nil {
			return
		}

		stmt, err := db.Prepare("UPDATE users SET name=? WHERE token=?")
		if err != nil {
			return
		}

		_, err = stmt.Exec(name, r.Header.Get("x-token"))
		if err != nil {
			return
		}

		fmt.Println(name)
	}
	return
}
