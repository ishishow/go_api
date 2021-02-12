package handler // 独自のクエリパッケージ

import (
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
	ctx := db.SetRepository(r.Context(), handler.Mysql)
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

	if err = usecase.Insert(ctx, &user); err != nil {
		return
	}
	service.RespondJSON(w, 200, user.Token)
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
	var err error
	user.Name, err = service.GainUserName(r)
	if err != nil {
		return
	}
	user.Token = r.Header.Get("x-token")
	user, err = usecase.Update(ctx, &user)

	service.RespondJSON(w, 200, user)
	return
}
