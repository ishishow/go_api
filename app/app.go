package app

import (
	"fmt"
	"net/http"

	"../db"      //実装した設定パッケージの読み込み
	"../handler" //実装したクエリパッケージの読み込み
)

func a() {

	db, err := db.ConnectDB()
	if err != nil {
		fmt.Println("error")
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
	} else {
		fmt.Println("データベース接続成功")
	}

	http.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateUser(w, r, db)
	})

	http.HandleFunc("/user/get", func(w http.ResponseWriter, r *http.Request) {
		handler.GetUser(w, r, db)
	})

	http.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		handler.UpdateUser(w, r, db)
	})

	http.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		handler.UpdateUser(w, r, db)
	})

	http.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		handler.UpdateUser(w, r, db)
	})

	http.ListenAndServe(":8080", nil)
}
