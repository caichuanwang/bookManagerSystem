package controller

import (
	"bookManagementSystem/api/feModal"
	"bookManagementSystem/modal"
	"bookManagementSystem/untils/sqlUntils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateRole(c echo.Context) error {
	r := modal.NewRole()
	if err := c.Bind(r); err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	createRoleStr := "insert into role(role_name,role_weight) values(?, ?)"
	stmt, err := db.Prepare(createRoleStr)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	_, err = stmt.Exec(r.Role_name, r.Role_weight)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("add success"))
}

func QueryRole(c echo.Context) error {
	whereCon := sqlUntils.CreateWhereSql(c, "role_name", "role_weight")
	orderBySql := sqlUntils.CreateOrderSql(c)
	LimitSql := sqlUntils.CreateLimitSql(c)
	queryRoleSql := fmt.Sprintf("select id,role_name,role_weight from role %s %s %s", whereCon, orderBySql, LimitSql)
	stmt, err := db.Prepare(queryRoleSql)
	if err != nil {
		return c.JSON(http.StatusOK, modal.Err(401, err.Error()))
	}
	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusOK, modal.Err(402, err.Error()))
	}
	defer rows.Close()
	var rs []feModal.Role
	for rows.Next() {
		var r modal.Role
		err := rows.Scan(&r.Id, &r.Role_name, &r.Role_weight.Int16)
		if err != nil {
			return c.JSON(http.StatusOK, modal.Err(404, err.Error()))
		}
		fer := modal.ToFeModal(r)
		rs = append(rs, *fer)
	}
	return c.JSON(http.StatusOK, modal.Success(rs))
}

func UpdateRole(c echo.Context) error {
	updateRoleSql := "update role set role_name = ? where id = ?"
	stmt, err := db.Prepare(updateRoleSql)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	_, err = stmt.Exec(c.FormValue("role_name"), c.FormValue("id"))
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("add success"))
}

func DeleteRole(c echo.Context) error {
	deleteRoleSql := "delete from role where id = ?"
	_, err := db.Exec(deleteRoleSql, c.QueryParam("id"))
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	return c.String(http.StatusOK, "success")
}
