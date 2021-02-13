package db

import (
	"testing"
	"time"

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

// func TestMysql_Insert(t *testing.T) {
// 	mysql := &Mysql{testdb.Setup()}
// 	defer mysql.Close()s

// 	user := &schema.User{
// 		ID:      nil,
// 		Name:    "ishishow",
// 		Token:   "uuuuid",
// 		Created: nil,
// 		Updated: nil,
// 	}

// 	err := Mysql.Insert(user)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestMysql_Get(t *testing.T) {
// 	Mysql := &Mysql{testdb.Setup()}
// 	defer Mysql.Close()

// 	token, err := service.CreateUuid()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	wantUser := &schema.User{
// 		ID:      nil,
// 		Name:    "ishishow",
// 		Token:   token,
// 		Created: nil,
// 		Updated: nil,
// 	}

// 	id, err := Mysql.Insert(wantUser)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	gotUser, err = Mysql.Get(token)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if equal(gotUser.Name, wantUser.Name) {
// 		t.Fatal("This is wrong user!")
// 	}
// }

// func equal(got interface{}, want interface{}) bool {
// 	return reflect.DeepEqual(got, want)
// }
