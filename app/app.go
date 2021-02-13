package app

import (
	"log"
	"net/http"
	"os"

	"../db"
	"github.com/gorilla/mux"
)

type App struct {
	router *mux.Router
	port   string
}

const defaultPort = "8080"

func (a *App) Initialize() {
	a.port = os.Getenv("PORT")
	if a.port == "" {
		a.port = defaultPort
	}
	var mysql *db.Mysql
	mysql, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("DB Initialize error: %v", err)
	}
	defer mysql.DB.Close()
	a.router = SetUpRouting(mysql)
	http.ListenAndServe(":"+a.port, a.router)

}

// func (a *App) Run() {
// 	http.ListenAndServe(":"+a.port, a.router)
// }
