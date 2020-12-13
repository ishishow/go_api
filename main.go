package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

const DRIVER = "mysql"
const DSN = "golang-test-user:golang-test-pass@tcp(mysql-container:3306)/golang-test-database"

func main() {
	db, err:= sql.Open(DRIVER, DSN)
	if err != nil {
		fmt.Println("Openエラー")
	} else {
		fmt.Println("OpenOK!")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("接続失敗！")
	} else {
		fmt.Println("接続OK!")
	}

	db.Close()
}