package schema

type Gacha struct {
	ID      int    `db:"ID"` //ID
	Name    string `db:"name"`
	Created string `db:"created_at"` //ID
	Updated string `db:"updated_at"` //ID
}
