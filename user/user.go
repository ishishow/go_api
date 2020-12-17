package user // 独自のクエリパッケージ

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"net/http"
	"io"
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
)


// マスタからSELECTしたデータをマッピングする構造体

type User struct {
	Id int `db:"ID"` //ID
	Name string `db:"name"` //ID
	Token string `db:"token"` //ID
	Created string `db:"created_at"` //ID
	Updated string `db:"updated_at"` //ID
}

func Create(w http.ResponseWriter, r *http.Request, db *sql.DB) (id int64, err error){
	switch r.Method {
		case "POST":
			var user User

			user.Name, err = GainUserName(r)
			if err != nil {
				return 0, err
			}

			user.Token, err = CreateUuid()
			if err != nil {
				return 0, err
			}

			stmt, err := db.Prepare("INSERT INTO users(name, token, created_at, updated_at) VALUES(?, ?, now(), now())")
			if err != nil {
				return 0, err
			}
			defer stmt.Close()

			//クエリ実行
			_, err = stmt.Exec(user.Name, user.Token)
			if err != nil {
				return 0, err
			}
	}
	return 0, nil
}


func Get(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case "GET":
		user, err := AuthUser(r.Header.Get("x-token"), db)
		if err != nil {
			return
		}
		fmt.Println(user.Name)
	}
	return
}

func Update(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// プリペアードステートメント
	switch r.Method {
		case "PUT":

			name, err := GainUserName(r)
			if err != nil {
				return
			}

			stmt, err := db.Prepare("UPDATE users SET name=? WHERE token=?")
			if err != nil {
				return
			}

			_, err = stmt.Exec(name, r.Header.Get("x-token"))
			if err != nil {
				return
			}

			fmt.Println(name)
	}
	return
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