package app

import (
	"database/sql"
	"fmt"
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
		fmt.Println("error")
	}
	defer mysql.DB.Close()
	a.router = SetUpRouting(mysql)
	checkDBHealth(mysql.DB)
	http.ListenAndServe(":"+a.port, a.router)

}

func checkDBHealth(targetDB *sql.DB) {
	if err := targetDB.Ping(); err != nil {
		fmt.Println("データベース接続失敗")
	} else {
		fmt.Println("データベース接続成功")
	}
}

// func (a *App) Run() {
// 	http.ListenAndServe(":"+a.port, a.router)
// }
