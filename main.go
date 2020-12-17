package main

import (
	"./db" //実装した設定パッケージの読み込み
	"./user" //実装したクエリパッケージの読み込み
	"fmt"
	"net/http"
)

func main() {

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
		user.Create(w, r, db)
	})

	http.HandleFunc("/user/get", func(w http.ResponseWriter, r *http.Request) {
		user.Get(w, r, db)
	})

	http.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		user.Update(w, r, db)
	})

	http.ListenAndServe(":8080", nil)

	//接続確認
	// // INSERTの実行
	// id, err := query.InsertUser("石川翔", "12345678", db)
	// if err != nil {
		// 	fmt.Println(err.Error())
		// }
	// fmt.Printf("登録されたユーザーのidは[%d]です。\n", id)

	// //SELECTの実行
	// user, err := query.SelectUserById(id, db)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

    // fmt.Printf("SELECTされたユーザ情報は以下の通りです。\n")
    // fmt.Printf("[ID] %s\n", user.Id)
    // fmt.Printf("[名前] %s\n", user.Name)
    // fmt.Printf("[token] %s\n", user.Token)
    // fmt.Printf("[登録日] %s\n", user.Created)
    // fmt.Printf("[登録日] %s\n", user.Updated)
}