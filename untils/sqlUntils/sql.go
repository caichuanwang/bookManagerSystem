package sqlUntils

import (
	"fmt"
	"github.com/samber/lo"
	"sort"
	"strconv"
	"strings"
)

func CreateWhereSql(param map[string]string, other ...string) string {
	if len(param) == 0 {
		return ""
	}
	var str = "where "
	var dbName = ""
	if other != nil && other[0] != "" {
		dbName = fmt.Sprintf("%s.", other[0])
	}
	keysArr := lo.Keys[string, string](param)
	sort.Strings(keysArr)

	lo.ForEach[string](keysArr, func(t string, i int) {
		if t != "" {
			str += fmt.Sprintf("%s%s LIKE '%%%s%%' and ", dbName, t, param[t])
		}
	})
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
