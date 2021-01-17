package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"../model"
	_ "github.com/go-sql-driver/mysql"
)

func AuthUser(token string, db *sql.DB) (user model.User, err error) {
	err = db.QueryRow("SELECT id, name FROM users WHERE token = ?", token).Scan(&user.ID, &user.Name)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("レコードが存在しません")
		return user, err
	case err != nil:
		return user, err
	default:
		fmt.Println(user.Name)
		return user, nil
	}
}

func GainUserName(r *http.Request) (name string, err error) {
	var user model.User
	body := r.Body
	defer body.Close()
	buf := new(bytes.Buffer)
	io.Copy(buf, body)

	// byte配列にしたbody内のjsonをgoで扱えるようにobjectに変換
	err = json.Unmarshal(buf.Bytes(), &user)
	if err != nil {
		fmt.Println("error 1")
		return "", err
	}
	fmt.Println(user.Name)
	return user.Name, nil
}
