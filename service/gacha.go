package service

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"../model"
	_ "github.com/go-sql-driver/mysql"
)

func GachaPlay(user model.User, times int, db *sql.DB) (gacha_draw_result GachaDrawRequest, err error) {

	gacha_draw_result, err = EmitCharacters(times, db)
	if err != nil {
		return gacha_draw_result, err
	}
	fmt.Println(gacha_draw_result)
	return gacha_draw_result, nil
}

func EmitCharacters(times int, db *sql.DB) (gacha_draw_result GachaDrawRequest, err error) {
	entries, sumWeight, err := SumWeight(db)
	if err != nil {
		return gacha_draw_result, err
	}
	var gacha_result GachaResult

	for i := 0; i < times; i++ {
		gacha_result.CharacterID, err = EmitCharacter(entries, sumWeight)
		if err != nil {
			return gacha_draw_result, err
		}

		gacha_result.Name, err = GetCharacterName(gacha_result.CharacterID, db)

		gacha_draw_result.Results = append(gacha_draw_result.Results, gacha_result)
	}
	return gacha_draw_result, nil
}

func EmitCharacter(entries []model.GachaEntries, sumWeight int) (emitCharacterID int, err error) {
	my_rand := rand.New(rand.NewSource(1))
	my_rand.Seed(time.Now().UnixNano())
	emitVal := my_rand.Intn(sumWeight)
	fmt.Println(my_rand.Intn(sumWeight))
	fmt.Println(my_rand.Intn(sumWeight))
	fmt.Println(my_rand.Intn(sumWeight))
	fmt.Println(my_rand.Intn(sumWeight))

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

	rows, err := db.Query("SELECT id, weight, character_id FROM gacha_entries WHERE gacha_id = ?", 1)
	if err != nil {
		return total_entries, sumWeight, err
	}

	for rows.Next() {
		entry := model.GachaEntries{}
		err = rows.Scan(&entry.ID, &entry.Weight, &entry.CharacterID)
		if err != nil {
			return total_entries, sumWeight, err
		}
		sumWeight += entry.Weight
		total_entries = append(total_entries, entry)
	}

	return total_entries, sumWeight, nil
}
