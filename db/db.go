package db

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

// DB設定の構造体
type ConfDB struct {
	Host string `json:"host"` //ホスト名
	Port int `json:"port"` //ポート番号
	DbName string `json:"db-name"` //接続先DB名
	Charset string `json:"charset"` //文字コード
	User string `json:"user"` //接続ユーザ名
	Pass string `json:"pass"` //接続パスワード
}

// URL設定の構造体
func ReadConfDB() (*ConfDB, error) {
	// 設定ファイル名
	const conffile = "db/db.json"
	// 構造体を初期化
	conf :=new(ConfDB)
	// 設定ファイルを読み込む
	cValue, err := ioutil.ReadFile(conffile)
	if err != nil {
		return conf, err
	}
	// 読み込んだjson文字列をデコードし、構造体にセット
	err = json.Unmarshal([]byte(cValue), conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}

func ConnectDB()(*sql.DB, error) {
	//設定ファイルを読み込む
	confDB, err := ReadConfDB()
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

	fmt.Println("ok")
	// // deferで終了前に必ずクローズする
	// defer db.Close()
	return db, err
}