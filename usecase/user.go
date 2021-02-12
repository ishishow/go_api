package service

import (
	"context"

	"../db"
	"../schema"
)

func Close(ctx context.Context) {
	db.Close(ctx)
}

func Get(ctx context.Context token string) (*schema.User, error) {
	return db.Get(ctx, token)
}

func Insert(ctx context.Context, todo *schema.Todo) (int, error) {
	return db.Insert(ctx, todo)
}

func Delete(ctx context.Context, id int) error {
	return db.Delete(ctx, id)
}

func GetAll(ctx context.Context) ([]schema.Todo, error) {
	return db.GetAll(ctx)
}
