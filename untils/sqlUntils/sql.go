package sqlUntils

import (
	"fmt"
	"strconv"
	"strings"
)

func CreateWhereSql(param map[string]interface{}) string {
	if len(param) == 0 {
		return ""
	}
	var str string = "where "
	for key, value := range param {
		if value != "" {
			str += fmt.Sprintf(" %s LIKE '%%%s%%' and ", key, value)
		}
	}
	if strings.LastIndex(str, "and") > 0 {
		return str[0:strings.LastIndex(str, "and")]
	}
	return ""
}

func CreateOrderSql(orderBy string, orderType string) string {
	if orderBy == "" {
		return ""
	}
	return fmt.Sprintf(" ORDER BY  %s %s", orderBy, orderType)
}

func CreateLimitSql(current int, pageSize int) string {
	return fmt.Sprintf(" LIMIT %s,%s", strconv.Itoa(pageSize*(current-1)), strconv.Itoa(pageSize))
}
