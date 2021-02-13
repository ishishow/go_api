package db

import (
	"context"

	"../schema"
)

const keyRepository = "Repository"

type Repository interface {
	Close() error
	Get(token string) (schema.User, error)
	Insert(User *schema.User) error
	Update(user *schema.User) (schema.User, error)
	GetAll() ([]schema.User, error)
}

func SetRepository(ctx context.Context, repository Repository) context.Context {
	return context.WithValue(ctx, keyRepository, repository)
}

func Close(ctx context.Context) error {
	return getRepository(ctx).Close()
}

func Get(ctx context.Context, token string) (schema.User, error) {
	return getRepository(ctx).Get(token)
}

func Insert(ctx context.Context, User *schema.User) error {
	return getRepository(ctx).Insert(User)
}

func Update(ctx context.Context, user *schema.User) (schema.User, error) {
	return getRepository(ctx).Update(user)
}

func GetAll(ctx context.Context) ([]schema.User, error) {
	return getRepository(ctx).GetAll()
}

func getRepository(ctx context.Context) Repository {
	return ctx.Value(keyRepository).(Repository)
}
