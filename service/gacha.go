package service

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"../schema"
	_ "github.com/go-sql-driver/mysql"
)

func GachaPlay(user schema.User, times int, db *sql.DB) (gacha_draw_result GachaDrawRequest, err error) {

	gacha_draw_result, err = EmitCharacters(times, db)
	if err != nil {
		return gacha_draw_result, err
	}
	err = SaveUserCharacter(user, gacha_draw_result, db)
	if err != nil {
		fmt.Println(err)
		return gacha_draw_result, err
	}
	return gacha_draw_result, nil
}

func SaveUserCharacter(user schema.User, gacha_draw_result GachaDrawRequest, db *sql.DB) (err error) {

	for _, result := range gacha_draw_result.Results {

		stmt, err := db.Prepare("INSERT INTO user_characters(character_id, user_id, created_at, updated_at) VALUES(?, ?, now(), now())")
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(result.CharacterID, user.ID)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil

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

func EmitCharacter(entries []schema.GachaEntries, sumWeight int) (emitCharacterID int, err error) {
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

func SumWeight(db *sql.DB) (total_entries []schema.GachaEntries, sumWeight int, err error) {
	sumWeight = 0

	rows, err := db.Query("SELECT id, weight, character_id FROM gacha_entries WHERE gacha_id = ?", 1)
	if err != nil {
		return total_entries, sumWeight, err
	}

	for rows.Next() {
		entry := schema.GachaEntries{}
		err = rows.Scan(&entry.ID, &entry.Weight, &entry.CharacterID)
		if err != nil {
			return total_entries, sumWeight, err
		}
		sumWeight += entry.Weight
		total_entries = append(total_entries, entry)
	}

	return total_entries, sumWeight, nil
}
