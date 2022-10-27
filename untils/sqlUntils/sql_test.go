package sqlUntils

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateWhereSql(t *testing.T) {
	Convey("should equal where noParams", t, func() {
		ass := assert.New(t)
		var params = map[string]string{"": ""}
		want := ""
		got := CreateWhereSql(params)
		So(got, ShouldEqual, want)
		ass.Equal(want, got)
		Convey("should equal where params", func() {
			params = map[string]string{
				"name": "Tom",
				"age":  "23",
				"time": "2022-5-27 15:13:14",
			}
			want := "where age LIKE '%23%' and name LIKE '%Tom%' and time LIKE '%2022-5-27 15:13:14%' "
			got := CreateWhereSql(params)
			So(got, ShouldEqual, want)
			Convey("should not equal where with dbName", func() {
				got := CreateWhereSql(params, "dbName")
				want := "where dbName.age LIKE '%23%' and dbName.name LIKE '%Tom%' and dbName.time LIKE '%2022-5-27 15:13:14%' "
				So(got, ShouldEqual, want)

			})
		})
	})
}
