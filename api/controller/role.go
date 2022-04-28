package controller

import (
	"bookManagerSystem/modal"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateRole @Summary 新增角色
// @Description 新增角色
// @Accept json
// @Param role body  modal.Role true "新增角色的数据"
// @Success 200 {object} modal.Result
// @Router /v1/role/add [post]
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

// QueryRole @Summary 查询角色
// @Description 查询角色
// @Accept json
// @Param req body modal.RoleParams false "查询角色的信息"
// @Success 200 {object} modal.RoleListResult
// @Router /v1/role/list [post]
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

// UpdateRole @Summary 修改角色
// @Description 修改角色
// @Accept json
// @Param role body  modal.Role true "新增角色的数据"
// @Success 200 {object} modal.Result
// @Router /v1/role/update [post]
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

// DeleteRole @Summary 删除角色
// @Description 删除角色
// @Accept json
// @Param id body string true "删除角色的id"
// @Success 200 {object} modal.Result
// @Router /v1/role/delete [delete]
func DeleteRole(c echo.Context) error {
	deleteRoleSql := "delete from role where id = ?"
	_, err := db.Exec(deleteRoleSql, c.QueryParam("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "success")
}

// QueryRoleOptions @Summary 查询角色的下拉选项
// @Description 查询角色的下拉选项
// @Accept json
// @Success 200 {object} modal.SelectOption
// @Router /v1/role/option [get]
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
