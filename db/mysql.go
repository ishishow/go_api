package db

import (
	"database/sql"

	"../schema"
)

type Mysql struct {
	DB *sql.DB
}

func (m *Mysql) Close() error {
	m.DB.Close()
	return nil
}

func (m *Mysql) Get(token string) (schema.User, error) {
	var user schema.User
	query := `SELECT id, name FROM users WHERE token = ?;`
	if err := m.DB.QueryRow(query, token).Scan(&user.ID, &user.Name); err != nil {
		return user, err
	}
	return user, nil
}

func (m *Mysql) Insert(user *schema.User) error {
	query := `INSERT INTO users(name, token, created_at, updated_at) VALUES(?, ?, now(), now());`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Name, user.Token); err != nil {
		return err
	}
	return nil
}

func (m *Mysql) Update(user *schema.User) (schema.User, error) {
	query := `UPDATE users SET name=? WHERE token=?;`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return *user, err
	}

	if _, err = stmt.Exec(user.Name, user.Token); err != nil {
		return *user, err
	}
	return *user, nil
}

func (m *Mysql) GetAll() ([]schema.User, error) {
	var UserList []schema.User
	return UserList, nil
}
