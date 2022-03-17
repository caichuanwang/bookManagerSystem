package sqlUntils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

func CreateWhereSql(c echo.Context, param ...string) string {
	if len(param) == 0 {
		return ""
	}
	var str string = "where "
	for _, item := range param {
		if c.FormValue("filter_"+item) != "" {
			str += fmt.Sprintf(" %s LIKE '%%%s%%' and ", item, c.FormValue("filter_"+item))
		}
	}
	return str[0:strings.LastIndex(str, "and")]
}

func CreateOrderSql(c echo.Context) string {
	var str string = " Order By "
	var val = c.FormValue("order_by")
	if val != "" {
		switch interface{}(val).(type) {
		case string:
			str += c.FormValue("order_by")
			break
		case []string:
			for _, item := range val {
				str += fmt.Sprintf(" %s ,", item)
			}
			break
		}
		return str
	} else {
		return ""
	}
}

func CreateLimitSql(c echo.Context) string {
	var str = " LIMIT "
	pageSize := c.FormValue("page_size")
	pageIndex := c.FormValue("page_index")
	if pageSize != "" && pageSize != "" {
		pageIndexInt, err := strconv.Atoi(pageIndex)
		pageSizeInt, err := strconv.Atoi(pageSize)
		if err != nil {
			return err.Error()
		}
		if pageSize != "" && pageIndex != "" {
			str += fmt.Sprintf("%s,%s", strconv.Itoa(pageSizeInt*(pageIndexInt-1)), pageSize)
			return str

		}
	}
	return ""
}
