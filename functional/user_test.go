// package functional

// import (
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/gorilla/mux"

// 	"../app"
// 	"../db"
// )

// func setupServer(mysql *db.Mysql) *mux.Router {
// 	return app.SetUpRouting(mysql)
// }

// func Test(t *testing.T) {
// 	mockdb, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("failed to init db mock")
// 	}
// 	defer mockdb.Close()
// 	mysqlMock := &db.Mysql{mockdb}

// 	router := setupServer(mysqlMock)
// 	testServer := httptest.NewServer(router)
// 	defer testServer.Close()

// }
