package sqlUntils

import (
	"fmt"
	"strconv"
	"strings"
)

func CreateWhereSql(param map[string]interface{}, other ...string) string {
	if len(param) == 0 {
		return ""
	}
	var str = "where "
	var dbName = ""
	if other != nil && other[0] != "" {
		dbName = fmt.Sprintf("%s.", other[0])
	}
	for key, value := range param {
		if value != "" {
			str += fmt.Sprintf(" %s%s LIKE '%%%s%%' and ", dbName, key, value)
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
