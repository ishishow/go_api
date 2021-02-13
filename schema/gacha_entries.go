package schema

type GachaEntries struct {
	ID          int    `db:"ID"`           //ID
	CharacterID int    `db:"character_id"` //character_id
	GachaID     int    `db:"gacha_id"`     //gacha_id
	Weight      int    `db:"weight"`       //weight
	Created     string `db:"created_at"`   //CreatedAt
	Updated     string `db:"updated_at"`   //UpdatedAt
}
