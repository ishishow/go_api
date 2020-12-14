package main

import (
	"./conf" //実装した設定パッケージの読み込み
	"database/sql" //実装した設定パッケージの読み込み
	"fmt" //実装した設定パッケージの読み込み
	_"github.com/go-sql-driver/mysql" //実装した設定パッケージの読み込み
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
}