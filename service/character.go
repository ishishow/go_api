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

func GetUserCharacters(user_id int, db *sql.DB) (list_user_character ListUserCharacterResponse, err error) {
	rows, err := db.Query("SELECT t1.id as UserCharacterID, t2.id as CharacterID, t2.name from user_characters AS t1 INNER JOIN characters AS t2 ON t1.character_id = t2.id AND t1.user_id = ?", user_id)
	if err != nil {
		return list_user_character, err
	}
	for rows.Next() {
		var user_character_response UserCharacterResponse
		err = rows.Scan(&user_character_response.UserCharacterID, &user_character_response.CharacterID, &user_character_response.Name)
		if err != nil {
			return list_user_character, err
		}
		list_user_character.Characters = append(list_user_character.Characters, user_character_response)
	}
	return list_user_character, nil

}
