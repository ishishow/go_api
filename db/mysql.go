package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../schema"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	DB *sql.DB
}

func (m *Mysql) Close() error {
	m.DB.Close()
	return nil
}

func (m *Mysql) Get(token string) (schema.User, error) {
	var user schema.User
	query := `SELECT id, name FROM USERS WHERE token = ?;`
	if err := m.DB.QueryRow(query, token).Scan(&user.ID, &user.Name, &user.Token, &user.Created, &user.Updated); err != nil {
		return user, err
	}
	return user, nil
}

func (m *Mysql) Insert(user *schema.User) error {
	query := `INSERT INTO users(name, token, created_at, updated_at) VALUES(?, ?, now(), now());`
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec(query, user.Name, user.Token); err != nil {
		return err
	}
	return nil
}

func (m *Mysql) Update(user *schema.User) (schema.User, error) {
	query := `UPDATE users SET name=? WHERE token=?;`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return *user, err
	}

	if _, err = stmt.Exec(user.Name, user.Token); err != nil {
		return *user, err
	}
	return *user, nil
}

func (m *Mysql) GetAll() ([]schema.User, error) {
	var UserList []schema.User
	return UserList, nil
}

type ConfDB struct {
	Host    string `json:"host"`    //ホスト名
	Port    int    `json:"port"`    //ポート番号
	DbName  string `json:"db-name"` //接続先DB名
	Charset string `json:"charset"` //文字コード
	User    string `json:"user"`    //接続ユーザ名
	Pass    string `json:"pass"`    //接続パスワード
}

// URL設定の構造体
func ReadConfDB() (*ConfDB, error) {
	// 設定ファイル名
	const conffile = "db/db.json"
	// 構造体を初期化
	conf := new(ConfDB)
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

func ConnectDB() (*Mysql, error) {
	confDB, err := ReadConfDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", confDB.User, confDB.Pass, confDB.Host, confDB.Port, confDB.DbName, confDB.Charset)
	db, err := sql.Open("mysql", conStr)
	if err != nil {
		fmt.Println(err.Error())
	}

	return &Mysql{db}, nil
}
