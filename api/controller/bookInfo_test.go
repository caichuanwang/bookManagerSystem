package controller

import (
	"bookManagerSystem/modal"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/smartystreets/goconvey/convey"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateBook(t *testing.T) {
	convey.Convey("all test begin", t, func() {
		e := echo.New()
		//sql mock
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		convey.Convey("create bookInfo where params correct", func() {
			var a = modal.CreateBookInfoParams{
				BookBaseInfo: modal.BookBaseInfo{Isbn: "2222", BookName: "test"},
				Photo:        []map[string]any{{"thumbUrl": "test", "name": "test"}},
			}
			str, _ := json.Marshal(a)
			req := httptest.NewRequest(echo.POST, "/v1/bookInfo/create", strings.NewReader(string(str)))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			mock.ExpectExec("INSERT INTO g_book_info").WithArgs("2222", "test").WillReturnResult(sqlmock.NewResult(1, 1))
			err = CreateBook(c)
			if err != nil {
				t.Errorf("expected error :%s", err.Error())
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
		})

	})
}
