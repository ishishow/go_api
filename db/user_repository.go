package db

import (
	"context"

	"../schema"
	"../service"
)

const keyRepository = "Repository"

type Repository interface {
	Close() error
	Get(token string) (schema.User, error)
	Insert(user *schema.User) error
	Update(user *schema.User) (schema.User, error)
	SumCharacterWeight() ([]schema.GachaEntries, int, error)
	GetCharacter(id int) (string , error)
	SaveUserCharacter(user_id int, gacha_draw_result service.GachaDrawRequest) error
}

func getRepository(ctx context.Context) Repository {
	return ctx.Value(keyRepository).(Repository)
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

func Insert(ctx context.Context, user *schema.User) error {
	return getRepository(ctx).Insert(user)
}

func Update(ctx context.Context, user *schema.User) (schema.User, error) {
	return getRepository(ctx).Update(user)
}

func (m *Mysql) Get(token string) (schema.User, error) {
	var user schema.User
	query := `SELECT * FROM users WHERE token = ?;`
	if err := m.DB.QueryRow(query, token).Scan(&user.ID, &user.Name, &user.Token, &user.Created, &user.Updated); err != nil {
		return user, err
	}
	return user, nil
}

func (m *Mysql) Insert(user *schema.User) error {
	query := `INSERT INTO users(name, token, created_at, updated_at) VALUES(?, ?, now(), now());`
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

	if _, err = tx.Exec(query, user.Name, user.Token); err != nil {
		return err
	}
	return nil
}

func (m *Mysql) Update(user *schema.User) (schema.User, error) {
	query := `UPDATE users SET name=? WHERE token=?;`
	tx, err := m.DB.Begin()
	if err != nil {
		return *user, err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec(query, user.Name, user.Token); err != nil {
		return *user, err
	}
	return *user, nil
}