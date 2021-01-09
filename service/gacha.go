package service

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"../model"
	_ "github.com/go-sql-driver/mysql"
)

func GachaPlay(user model.User, str_times string, db *sql.DB) (err error) {

	//
	sumWeight, err := SumWeight(db)
	times, err := strconv.Atoi(str_times)
	characters, err := EmitCharacters(times, sumWeight, db)

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

func EmitCharacters(times int, sumWeight int, db *sql.DB) (characters_id []int, err error) {

	rand.Seed(time.Now().UnixNano())
	emitVal := rand.Intn(sumWeight)

	return characters_id, nil
}

func SumWeight(db *sql.DB) (sumWeight int, err error) {
	sumWeight = 0

	rows, err := db.Query("SELECT weight FROM gacha_entries WHERE gacha_id = ?", 1)
	if err != nil {
		return sumWeight, err
	}

	for rows.Next() {
		m := model.GachaEntries{}
		err = rows.Scan(&m.Weight)
		if err != nil {
			return sumWeight, err
		}
		sumWeight += m.Weight
	}

	return sumWeight, nil
}
