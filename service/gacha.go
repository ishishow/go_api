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

func GachaPlay(user model.User, str_times string, db *sql.DB) (gacha_draw_result model.GachaDrawRequest, err error) {

	times, err := strconv.Atoi(str_times)
	if err != nil {
		return nil, err
	}
	gacha_draw_result, err = EmitCharacters(times, db)
	if err != nil {
		return gacha_draw_result, err
	}

	return gacha_draw_result, nil
}

func EmitCharacters(times int, db *sql.DB) (gacha_draw_result model.GachaDrawRequest, err error) {
	entries, sumWeight, err := SumWeight(db)
	if err != nil {
		return nil, err
	}
	var gacha_result GachaResult

	for i, _ := range times {
		gacha_result.CharacterID, err = EmitCharacter(entries, sumWeight)
		if err != nil {
			return gacha_draw_result, err
		}

		gacha_result.Name, err = GetCharacterName(emitedCharacterID, db)


		gacha_draw_result = append(gacha_draw_result, gacha_result)
	}
	return gacha_draw_result, nil
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