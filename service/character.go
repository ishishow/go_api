package service

import (
	"database/sql"
	"fmt"

	"../model"
)

func GetCharacterName(character_id int, db *sql.DB) (name string, err error) {

	var character model.Chatacter
	err = db.QueryRow("SELECT name FROM characters WHERE id = ?", character_id).Scan(&character.Name)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("レコードが存在しません。 character_id = ", character_id)
		return character.Name, err
	case err != nil:
		panic(err.Error())
		return character.Name, err
	default:
		return character.Name, nil
	}
}
