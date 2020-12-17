package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"io"
	"bytes"
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
	_"../model"
)

type User struct {
	Id int `db:"ID"` //ID
	Name string `db:"name"` //ID
	Token string `db:"token"` //ID
	Created string `db:"created_at"` //ID
	Updated string `db:"updated_at"` //ID
}

func CreateUuid()(token string, err error){
	u, err := uuid.NewRandom()
	if err != nil {
			fmt.Println(err)
			return
	}
	uu := u.String()
	return uu, err
}

func AuthUser(token string, db *sql.DB)(user User, err error){
	err = db.QueryRow("SELECT name FROM users WHERE token = ?", token).Scan(&user.Name)
	switch {
		case err == sql.ErrNoRows:
			fmt.Println("レコードが存在しません")
			return user, err
		case err != nil:
			panic(err.Error())
			return user, err
		default:
			fmt.Println(user.Name)
		return user, nil
	}
}

func GainUserName(r *http.Request)(name string, err error){
	var user User
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