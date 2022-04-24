package controller

import (
	"bookManagerSystem/modal"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateRole(c echo.Context) error {
	r := modal.NewRole()
	if err := c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	createRoleStr := "insert into role(role_name,role_weight) values(?, ?)"
	stmt, err := db.Prepare(createRoleStr)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	_, err = stmt.Exec(r.Role_name, r.Role_weight)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("add success"))
}

func QueryRole(c echo.Context) error {
	queryRoleSql := fmt.Sprintf("select id,role_name,role_weight from role")
	stmt, err := db.Prepare(queryRoleSql)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	var rs []modal.RoleListResult
	for rows.Next() {
		var r modal.RoleListResult
		err := rows.Scan(&r.Id, &r.Role_name, &r.Role_weight)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		rs = append(rs, r)
	}
	queryCount := "select COUNT(id) from role"
	var a int
	db.QueryRow(queryCount).Scan(&a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.TableSucc(rs, a))
}

func UpdateRole(c echo.Context) error {
	r := modal.NewRole()
	if err := c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	updateRoleSql := "update role set role_name = ? where id = ?"
	stmt, err := db.Prepare(updateRoleSql)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	_, err = stmt.Exec(r.Role_name, r.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("add success"))
}

func DeleteRole(c echo.Context) error {
	deleteRoleSql := "delete from role where id = ?"
	_, err := db.Exec(deleteRoleSql, c.QueryParam("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "success")
}

func QueryRoleOptions(c echo.Context) error {
	queryRole := "select id,role_name from role"
	rows, err := db.Query(queryRole)
	if err != nil {
		return err
	}
	cols, err := rows.Columns()
	if err != nil {
		return err
	}
	vals := make([]sql.RawBytes, len(cols)) //建立接口 [id role_name]
	valsp := make([]interface{}, len(vals)) //建立接口指针的接口
	result := make([]modal.SelectOption, 0)
	//将接口转换为指针类型的接口
	for i := range vals {
		valsp[i] = &vals[i]
	}
	// valps [&id,&role_name]
	//解析查询结果
	for rows.Next() {
		if err := rows.Scan(valsp...); err == nil { //注意：此处用valsp
			//var value string
			//for i, col := range vals { //注意：此处用vals
			//	if col == nil {
			//		value = "NULL"
			//	} else {
			//		value = string(col)
			//	}
			//	//注意：读取的数据是uint8类型的数组，需要转成byte类型的数组才好转换成其他
			//	fmt.Println(cols[i], ":", v salue) //输出各个列
			//}
		} else {
			return err
		}
		result = append(result, modal.SelectOption{
			Label: string(vals[1]),
			Value: string(vals[0]),
		})
	}
	return c.JSON(http.StatusOK, result)
}
