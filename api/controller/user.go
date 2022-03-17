package controller

import (
	"bookManagementSystem/modal"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateAddUser 添加用户
func CreateAddUser(c echo.Context) error {
	var u = modal.NewUser()
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	createUserStr := "insert into user(user_name,user_password,sex,birthday,borrow_book_count,phone,email,remake,role) values(?,?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(createUserStr)
	if err != nil {
		return c.String(http.StatusOK, err.Error())

	}
	defer stmt.Close()
	Result, err := stmt.Exec(u.User_name, u.User_password, u.Sex, u.Birthday, u.Borrow_book_count, u.Phone, u.Email, u.Remake, u.Role)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	fmt.Println(Result.RowsAffected())
	return c.JSON(http.StatusOK, modal.Success("add success"))
}

func QueryUser(c echo.Context) error {
	fmt.Println(c.Request().Header)
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	fmt.Println(name)
	u := modal.NewUserWithAllKeys()
	queryUserStr := "select user_name,user_password,sex,birthday,borrow_book_count,phone,email,remake,role from user"

	if userName := c.FormValue("filter_user_name"); userName != "" {
		queryUserStr = queryUserStr + " where user_name= " + fmt.Sprintf("'%s'", userName)
	}
	if sex := c.FormValue("order_by"); sex != "" {
		queryUserStr = queryUserStr + "order by " + sex
	}
	stmt, err := db.Prepare(queryUserStr)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow().Scan(&u.User_name, &u.User_password, &u.Sex, &u.Birthday, &u.Borrow_book_count, &u.Phone, &u.Remake, &u.Email, &u.Role)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success(&u))
}

func UpdateUser(c echo.Context) error {
	updateUserSQL := "update user set user_name = ? where id = ?"
	stmt, err := db.Prepare(updateUserSQL)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	_, err = stmt.Exec(c.FormValue("user_name"), c.FormValue("id"))
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("update success"))
}
func DeleteUser(c echo.Context) error {
	deleteUserSql := "delete from user where id = ?"
	fmt.Println(c.QueryParam("id"))
	stmt, err := db.Prepare(deleteUserSql)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	_, err = stmt.Exec(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("delete success"))
}
