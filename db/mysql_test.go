package db

import (
	"testing"
	"time"

	"../schema"
	"../service"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestSQLMock_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to init db mock")
	}
	defer db.Close()
	mysqlMock := &Mysql{db}

	token, err := service.CreateUuid()
	if err != nil {
		t.Fatalf("uuid error")
	}

	columns := []string{"id", "token", "name", "created_at", "updated_at"}
	mock.ExpectQuery("SELECT (.+) FROM USERS WHERE token =").
		WithArgs(token).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "ishishow", token, time.Now(), time.Now()))

	if _, err := mysqlMock.Get(token); err != nil {
		t.Fatalf("failed to get user: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("failed to ExpectationWereMet(): %s", err)
	}
}

func TestSQLMock_Insert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to init db mock")
	}
	defer db.Close()
	mysqlMock := &Mysql{db}

	token, err := service.CreateUuid()
	if err != nil {
		t.Fatalf("uuid error")
	}

	user := &schema.User{Name: "ishishow", Token: token}
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO users`).
		WithArgs("ishishow", token).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err := mysqlMock.Insert(user); err != nil {
		t.Fatalf("failed to insert user: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("failed to ExpectationWereMet(): %s", err)
	}
}

func TestSQLMock_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to init db mock")
	}
	defer db.Close()
	mysqlMock := &Mysql{db}

	token, err := service.CreateUuid()
	if err != nil {
		t.Fatalf("uuid error")
	}

	user := &schema.User{Name: "ishishow", Token: token}
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE users`).
		WithArgs("ishishow", token).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if _, err := mysqlMock.Update(user); err != nil {
		t.Fatalf("failed to update user: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("failed to ExpectationWereMet(): %s", err)
	}
}
