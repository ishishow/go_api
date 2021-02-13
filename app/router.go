package app

import (
	"net/http"

	"../db"
	"../handler"
	"github.com/gorilla/mux"
)

type route struct {
	method      string
	path        string
	handlerFunc http.HandlerFunc
}

func SetUpRouting(mysql *db.Mysql) *mux.Router {

	UserHandler := &handler.UserHandler{
		Mysql: mysql,
	}

	routes := []route{
		route{"GET", "/user/get", UserHandler.GetUser},
		route{"POST", "/user/create", UserHandler.CreateUser},
		route{"PUT", "/user/update", UserHandler.UpdateUser},
		route{"POST", "/gacha/draw", UserHandler.GachaDraw},
		route{"GET", "/character/list", UserHandler.GetUserCharacterAll},
	}

	router := mux.NewRouter()
	for _, route := range routes {
		router.Methods(route.method).Path(route.path).Handler(route.handlerFunc)
	}

	return router
}
