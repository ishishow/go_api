package handler // 独自のクエリパッケージ

import (
	"fmt"
	"net/http"

	"../db"
	"../schema"
	"../service"
	"../usecase"
	_ "github.com/go-sql-driver/mysql"
)

// マスタからSELECTしたデータをマッピングする構造体

type UserHandler struct {
	Mysql *db.Mysql
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user schema.User
	var err error
	user.Name, err = service.GainUserName(r)
	if err != nil {
		return
	}

	user.Token, err = service.CreateUuid()
	if err != nil {
		return
	}

	stmt, err := handler.Mysql.DB.Prepare("INSERT INTO users(name, token, created_at, updated_at) VALUES(?, ?, now(), now())")
	if err != nil {
		return
	}
	defer stmt.Close()

	//クエリ実行
	_, err = stmt.Exec(user.Name, user.Token)
	if err != nil {
		return
	}

	return
}

func (handler *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handler.Mysql)
	user, err := usecase.Get(ctx, r.Header.Get("x-token"))
	if err != nil {
		return
	}
	service.RespondJSON(w, 200, user)
	return
}

func (handler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handler.Mysql)

	var user schema.User
	user.Name, err := service.GainUserName(r)
	if err != nil {
		return
	}
	user.Token = r.Header.Get("x-token")
	user, err = usecase.Update(ctx, &user)

	stmt, err := handler.Mysql.DB.Prepare("UPDATE users SET name=? WHERE token=?")
	if err != nil {
		return
	}

	_, err = stmt.Exec(name, r.Header.Get("x-token"))
	if err != nil {
		return
	}

	fmt.Println(name)
	return
}
