package service

import (

	"database/sql"
	"../schema"
	_ "github.com/go-sql-driver/mysql"
)

func AuthUser(token string, db *sql.DB) (user schema.User, err error) {
	err = db.QueryRow("SELECT id, name FROM users WHERE token = ?", token).Scan(&user.ID, &user.Name)
	switch {
	case err == sql.ErrNoRows:
		return user, err
	case err != nil:
		return user, err
	default:
		return user, nil
	}
}

// func GainUserName(r *http.Request) (name string, err error) {
// 	var user schema.User
// 	body := r.Body
// 	defer body.Close()
// 	buf := new(bytes.Buffer)
// 	io.Copy(buf, body)

// 	if err := json.Unmarshal(buf.Bytes(), &user); err != nil {
// 		return "", err
// 	}
// 	return user.Name, nil
// }
