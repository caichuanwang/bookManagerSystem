package controller

import (
	"bookManagerSystem/modal"
	"bookManagerSystem/untils/sqlUntils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary 查询用户
// @Description 查询用户
// @Accept json
// @Param user_name body string true "userName"
// @Param user_password body string true "userPassword"
// @Success 200 {object} modal.Result
// @Router /user/add [post]
func CreateAddUser(c echo.Context) error {
	var u = modal.NewUser()
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	ok, err := govalidator.ValidateStruct(u)
	if err != nil || !ok {
		return c.String(http.StatusBadRequest, err.Error())
	}
	createUserStr := "insert into user(user_name,user_password,sex,birthday,borrow_book_count,phone,email,remake,role) values(?,?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(createUserStr)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())

	}
	defer stmt.Close()
	Result, err := stmt.Exec(u.User_name, u.User_password, u.Sex, u.Birthday, u.Borrow_book_count, u.Phone, u.Email, u.Remake, u.Role)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(Result.RowsAffected())
	return c.JSON(http.StatusOK, modal.Success("add success"))
}

func QueryUser(c echo.Context) error {
	var u = new(modal.QueryUserParams)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	var paramMap = make(map[string]interface{})
	paramMap["user_name"] = u.User_name
	paramMap["role"] = u.Role
	paramMap["borrow_book_count"] = u.Borrow_book_count
	whereCon := sqlUntils.CreateWhereSql(paramMap)
	orderBySql := sqlUntils.CreateOrderSql(u.Order_by, u.Order_type)
	LimitSql := sqlUntils.CreateLimitSql(u.Current, u.PageSize)
	queryUserStr := fmt.Sprintf("select id,user_name,sex,birthday,borrow_book_count,phone,email,remake,role, (SELECT role_name FROM role WHERE role.id = role ) AS roleName from user  %s %s %s", whereCon, orderBySql, LimitSql)
	stmt, err := db.Prepare(queryUserStr)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	var rs []modal.UserListResult
	for rows.Next() {
		var r modal.UserListResult
		err := rows.Scan(&r.Id, &r.User_name, &r.Sex, &r.Birthday, &r.Borrow_book_count, &r.Phone, &r.Email, &r.Remake, &r.Role, &r.RoleName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		rs = append(rs, r)
	}
	queryCount := "select COUNT(id) from user"
	var a int
	db.QueryRow(queryCount).Scan(&a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.TableSucc(rs, a))
}

func UpdateUser(c echo.Context) error {
	var u = modal.NewUser()
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	ok, err := govalidator.ValidateStruct(u)
	if err != nil || !ok {
		return c.String(http.StatusBadRequest, err.Error())
	}
	updateUserSQL := "update user set user_name = ?,role = ? ,sex = ? ,birthday = ?,email=?,phone = ?,remake=?,borrow_book_count=? where id = ?"
	stmt, err := db.Prepare(updateUserSQL)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	_, err = stmt.Exec(&u.User_name, &u.Role, &u.Sex, &u.Birthday, &u.Email, &u.Phone, &u.Remake, &u.Borrow_book_count, &u.Id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("update success"))
}
func DeleteUser(c echo.Context) error {
	deleteUserSql := "delete from user where id = ?"
	stmt, err := db.Prepare(deleteUserSql)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	_, err = stmt.Exec(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("delete success"))
}
