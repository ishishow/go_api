package model

type Chatacter struct {
	Id int `db:"ID"` //ID
	Name string `db:"name"` //ID
	Rarity int `db:"rarity"` //ID
	Created string `db:"created_at"` //ID
	Updated string `db:"updated_at"` //ID
}