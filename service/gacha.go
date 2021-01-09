package service

import (
	"database/sql"
	"fmt"
	"strconv"

	"../model"
	_ "github.com/go-sql-driver/mysql"
)

func GachaPlay(user model.User, str_times string, db *sql.DB) (err error) {

	//
	times, err := strconv.Atoi(str_times)
	characters, err := EmitCharacters(times, db)

	err = db.QueryRow("SELECT name FROM users WHERE token = ?", token).Scan(&user.Name)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("レコードが存在しません")
		return err
	case err != nil:
		panic(err.Error())
		return err
	default:
		fmt.Println(user.Name)
		return nil
	}
}

func EmitCharacters(times int, db *sql.DB) (characters_id []int, err error) {

}

func SumWeight(db *sql.DB) (sum_weight int, err error) {
	rows, err := db.Query("SELECT weight FROM gacha_entries WHERE gacha_id = ?", 1)
	if err != nil {
		return
	}

	for rows.Next() {
		m := model.GachaEntries{}
		err = rows.Scan(&m.ID, &m.Weight)
		if err != nil {
			return
		}
		members = append(members, m)
	}
}
