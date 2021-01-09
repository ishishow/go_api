package model

type UserCharacter struct {
	ID          int    `db:"ID"`           //ID
	CharacterID int    `db:"character_id"` //ID
	UserID      int    `db:"user_id"`      //ID
	Created     string `db:"created_at"`   //ID
	Updated     string `db:"updated_at"`   //ID
}
