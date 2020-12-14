package main

import (
	"./conf" //実装した設定パッケージの読み込み
	"./query" //実装したクエリパッケージの読み込み
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func main() {

	//設定ファイルを読み込む
	confDB, err := conf.ReadConfDB()
	if err != nil {
		fmt.Println(err.Error())
	}

	//設定から接続文字列を生成
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", confDB.User, confDB.Pass, confDB.Host, confDB.Port, confDB.DbName, confDB.Charset)

	// データベース接続
	db, err := sql.Open("mysql", conStr)
	if err != nil {
		fmt.Println(err.Error())
	}

	// deferで終了前に必ずクローズする
	defer db.Close()

	//接続確認
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
	} else {
		fmt.Println("データベース接続成功")
	}

	// INSERTの実行
	id, err := query.InsertUser("石川翔", "12345678", db)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("登録されたユーザーのidは[%d]です。\n", id)

	//SELECTの実行
	user, err := query.SelectUserById(id, db)
	if err != nil {
		fmt.Println(err.Error())
	}

    fmt.Printf("SELECTされたユーザ情報は以下の通りです。\n")
    fmt.Printf("[ID] %s\n", user.Id)
    fmt.Printf("[名前] %s\n", user.Name)
    fmt.Printf("[token] %s\n", user.Token)
    fmt.Printf("[登録日] %s\n", user.Created)
    fmt.Printf("[登録日] %s\n", user.Updated)
}