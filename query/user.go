package query // 独自のクエリパッケージ

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


// マスタからSELECTしたデータをマッピングする構造体

type Users struct {
	Id int `db:"ID"` //ID
	Name string `db:"name"` //ID
	Token string `db:"token"` //ID
	Created string `db:"created_at"` //ID
	Updated string `db:"updated_at"` //ID
}

// データ登録関数
func InsertUser(name, token string, db *sql.DB)(id int64, err error) {

	// プリペアードステートメント
	stmt, err := db.Prepare("INSERT INTO users(name, token, created_at, updated_at) VALUES(?, ?, now(), now())")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	//クエリ実行
	result, err := stmt.Exec(name, token)
	if err != nil {
		return 0, err
	}

	// オートインクリメントのIDを取得
	insertedId, err  := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	//INSERTしたIDを返す
	return insertedId, nil

}


func SelectUserById(id int64, db *sql.DB)(userinfo Users, err error) {

	//構造体Users型の変数userを宣言
	var user Users

	//プリペアードステートメント
	stmt, err := db.Prepare("SELECT ID, name, token, created_at, updated_at FROM users WHERE ID = ?")
	if err != nil {
		return user, err
	}

	//クエリ実行
	rows, err := stmt.Query(id)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	//SELECTした結果を構造体にマップ
	rows.Next()
	err = rows.Scan(&user.Id, &user.Name, &user.Token, &user.Created, &user.Updated)
	if err != nil {
		return user, err
	}

	// 取得データをマッピングしたUSER構造体を返す
	return user, nil

}


// 全行SELECT用の構造体配列
type UserList []Users

// 全行データ取得関数
func SelectUserAll(db *sql.DB) (userlist UserList, err error) {

    // 配列宣言
    var ul UserList

    // プリペアードステートメント
    stmt, err := db.Prepare("SELECT ID,name,token,created_at,updated_at FROM USERS")
    if err != nil {
        return ul, err
    }

    // クエリ実行
    rows, err := stmt.Query()
    if err != nil {
        return ul, err
    }
    defer rows.Close()

    // SELECTした結果を構造体にマップ
    for rows.Next() {
        // 構造体宣言
        var user Users
        err = rows.Scan(&user.Id, &user.Name, &user.Token, &user.Created, &user.Updated)
        // 配列にScan結果を追加
        ul = append(ul, user)
    }

    // 取得データをマッピングしたM_USER構造体配列を返す
    return ul, nil
}