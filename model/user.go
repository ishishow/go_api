package model

type User struct {
	Id int `db:"ID"` //ID
	Name string `db:"name"` //ID
	Token string `db:"token"` //ID
	Created string `db:"created_at"` //ID
	Updated string `db:"updated_at"` //ID
}