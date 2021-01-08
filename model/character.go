package model

type Chatacter struct {
	Id int `db:"ID"` //ID
	Name string `db:"name"` //ID
	Created string `db:"created_at"` //ID
	Updated string `db:"updated_at"` //ID
}