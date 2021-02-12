package usecase

import (
	"context"

	"../db"
	"../schema"
)

func Close(ctx context.Context) error {
	return db.Close(ctx)
}

func Get(ctx context.Context, token string) (schema.User, error) {
	return db.Get(ctx, token)
}

func Insert(ctx context.Context, User *schema.User) (int, error) {
	return db.Insert(ctx, User)
}

func Update(ctx context.Context, user *schema.User) (schema.User, error) {
	return db.Update(ctx, user)
}

func GetAll(ctx context.Context) ([]schema.User, error) {
	return db.GetAll(ctx)
}
