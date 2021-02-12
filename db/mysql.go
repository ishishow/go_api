package db

import (
	"database/sql"

	"../schema"
)

type Mysql struct {
	DB *sql.DB
}

func (m *Mysql) Close() {
	m.DB.Close()
}

func (m *Mysql) Get(token string) (schema.User, error) {
	var user schema.User
	query := `SELECT id, name FROM users WHERE token = ?;`
	if err := m.DB.QueryRow(query, token).Scan(&user.ID, &user.Name); err != nil {
		return user, err
	}
	return user, nil
}

// func (p *Postgres) Insert(User *schema.User) (int, error) {
// 	query := `
//         INSERT INTO User (id, title, note, due_date)
//         VALUES (nextval('User_id'), $1, $2, $3)
//         RETURNING id;
//     `

// 	rows, err := p.DB.Query(query, User.Title, User.Note, User.DueDate)
// 	if err != nil {
// 		return -1, err
// 	}

// 	var id int
// 	for rows.Next() {
// 		if err := rows.Scan(&id); err != nil {
// 			return -1, err
// 		}
// 	}

// 	return id, nil
// }

// func (p *Postgres) GetAll() ([]schema.User, error) {
// 	query := `
//         SELECT *
//         FROM User
//         ORDER BY id;
//     `

// 	rows, err := p.DB.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var UserList []schema.User
// 	for rows.Next() {
// 		var t schema.User
// 		if err := rows.Scan(&t.ID, &t.Title, &t.Note, &t.DueDate); err != nil {
// 			return nil, err
// 		}
// 		UserList = append(UserList, t)
// 	}

// 	return UserList, nil
// }

// func ConnectPostgres() (*Postgres, error) {
// 	connStr := "postgres://postgres:postgres@postgres/postgres?sslmode=disable"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Postgres{db}, nil
// }
