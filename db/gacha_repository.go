package db

import (
	"context"
	"../schema"
	"../service"

)

func SumCharacterWeight(ctx context.Context) ([]schema.GachaEntries, int, error) {
	return getRepository(ctx).SumCharacterWeight()
}

func SaveUserCharacter(ctx context.Context, user_id int, gacha_draw_request service.GachaDrawRequest) error {
	return getRepository(ctx).SaveUserCharacter(user_id, gacha_draw_request)
}

func GetCharacter(ctx context.Context, id int) (string, error) {
	return getRepository(ctx).GetCharacter(id)
}
func (m *Mysql) SumCharacterWeight() ([]schema.GachaEntries, int, error) {
	var total_entries []schema.GachaEntries
	sumWeight := 0
	query := `SELECT id, weight, character_id FROM gacha_entries WHERE gacha_id = ?;`

	rows, err := m.DB.Query(query, 1)
	if err != nil {
		return total_entries, sumWeight, err
	}

	for rows.Next() {
		entry := schema.GachaEntries{}
		err = rows.Scan(&entry.ID, &entry.Weight, &entry.CharacterID)
		if err != nil {
			return total_entries, sumWeight, err
		}
		sumWeight += entry.Weight
		total_entries = append(total_entries, entry)
	}

	return total_entries, sumWeight, nil
}


func (m *Mysql) GetCharacter(character_id int) (string, error) {
	query := `SELECT name FROM characters WHERE id = ?;`
	var character schema.Chatacter
	if err := m.DB.QueryRow(query, character_id).Scan(&character.Name); err != nil {
		return "", err
	} 
	return character.Name, nil
}

func (m *Mysql) SaveUserCharacter(user_id int, gacha_draw_result service.GachaDrawRequest) error {
	query := `INSERT INTO user_characters(character_id, user_id, created_at, updated_at) VALUES(?, ?, now(), now());`
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()
	
	for _, result := range gacha_draw_result.Results {
		if _, err = tx.Exec(query, result.CharacterID, user_id); err != nil {
			return err
		}	
	}
	return nil
}