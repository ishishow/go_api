package service

import (
	"database/sql"
	"fmt"

	"../model"
)

type ListUserCharacterResponse struct {
	Characters []UserCharacterResponse
}

type UserCharacterResponse struct {
	UserCharacterID string `json:"user_character_id"`
	CharacterID     string `json:"character_id"`
	Name            string `json:"name"`
}

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

func GerUserCharacters(user_id int, db *sql.DB) (list_user_character ListUserCharacterResponse, err error) {

	rows, err := db.Query("SELECT  ")
}
