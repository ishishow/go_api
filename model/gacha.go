package model

type Gacha struct {
	Id int `db:"ID"` //ID
	CharacterId int `db:"character_id"` //ID
	Weight int `db:"weight"` //ID
	Created string `db:"created_at"` //ID
	Updated string `db:"updated_at"` //ID
}