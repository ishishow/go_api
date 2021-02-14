package usecase

import (
	"context"

	"../db"
	"../service"
)

func PlayGacha(ctx context.Context, user_id int, times int) (service.GachaDrawRequest, error) {
	var gacha_draw_result service.GachaDrawRequest
	gacha_draw_result, err := EmitCharacters(ctx, times)
	if err != nil {
		return gacha_draw_result, err
	}

	if err := db.SaveUserCharacter(ctx, user_id, gacha_draw_result); err != nil {
		return gacha_draw_result, err
	}

	return gacha_draw_result, nil
}


func EmitCharacters(ctx context.Context, times int) (service.GachaDrawRequest, error) {
	var gacha_draw_result service.GachaDrawRequest
	entries, sumWeight, err := db.SumCharacterWeight(ctx)
	if err != nil {
		return gacha_draw_result, err
	}
	var gacha_result service.GachaResult

	for i := 0; i < times; i++ {
		gacha_result.CharacterID, err = service.EmitCharacter(entries, sumWeight)
		if err != nil {
			return gacha_draw_result, err
		}

		gacha_result.Name, err = db.GetCharacter(ctx, gacha_result.CharacterID)
		gacha_draw_result.Results = append(gacha_draw_result.Results, gacha_result)
	}
	return gacha_draw_result, nil
}


// entries, sumWeight, err := db.SumWeight()


// func Update(ctx context.Context, user *schema.User) (schema.User, error) {
// 	return db.Update(ctx, user)
// }