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

func GachaPlay(user model.User, str_times string, db *sql.DB) (character_ids []int, err error) {

	times, err := strconv.Atoi(str_times)
	if err != nil {
		return nil, err
	}
	character_ids, err = EmitCharacters(times, db)
	if err != nil {
		return character_ids, err
	}

	return character_ids, nil
}

func EmitCharacters(times int, db *sql.DB) (character_ids []int, err error) {
	entries, sumWeight, err := SumWeight(db)
	if err != nil {
		return character_ids, err
	}

	for i, _ := range times {
		emitedCharacterID, err := EmitCharacter(entries, sumWeight)
		if err != nil {
			return character_ids, err
		}
		character_ids = append(character_ids, emitedCharacterID)
	}
	return character_ids, nil
}

func EmitCharacter(entries []model.GachaEntries, sumWeight int) (emitCharacterID int, err error {
	rand.Seed(time.Now().UnixNano())
	emitVal := rand.Intn(sumWeight)

	for _, entry := range entries {
		emitVal -= entry.Weight
		if emitVal <= 0 {
			emitCharacterID = entry.CharacterID
			return emitCharacterID, nil
		}
	}
	return 0, err
}

func SumWeight(db *sql.DB) (total_entries []model.GachaEntries, sumWeight int, err error) {
	sumWeight = 0

	rows, err := db.Query("SELECT weight FROM gacha_entries WHERE gacha_id = ?", 1)
	if err != nil {
		return sumWeight, err
	}

	for rows.Next() {
		entry := model.GachaEntries{}
		err = rows.Scan(&entry.ID, &entry.Weight, &entry.CharacterID)
		if err != nil {
			return nil, sumWeight, err
		}
		sumWeight += entry.Weight
		total_entries = append(total_entries, entry)
	}

	return total_entries, sumWeight, nil
}

func UserCharacterCreate(userID int, character_ids []int)(err error){

	for _, character_id := range character_ids{

		stmt, err := db.Prepare("INSERT INTO user_characters(user_id, character_id, created_at, updated_at) VALUES(?, ?, now(), now())")
		if err != nil {
			return 0, err
		}
		defer stmt.Close()

		//クエリ実行
		_, err = stmt.Exec(userID, character_id)
		if err != nil {
			return 0, err
		}
	}

}