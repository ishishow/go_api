package model

type UserCharacter struct {
	Id int `db:"ID"` //ID
	CharacterId int `db:"character_id"` //ID
	UserId int `db:"user_id"` //ID
	Created string `db:"created_at"` //ID
	Updated string `db:"updated_at"` //ID
}