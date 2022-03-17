package controller

import (
	"bookManagementSystem/api/middleware"
	"bookManagementSystem/modal"
	"bookManagementSystem/untils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"net/http"
)

var database = untils.ReadCon("mysql", "databaseName")

func HandleLoginController(c echo.Context) error {
	username := c.FormValue("user_name")
	password := c.FormValue("password")
	var u = modal.NewUser()
	queryStr := fmt.Sprintf("select user_name,user_password,role from %s where user_name = ?", "user")
	row := db.QueryRow(queryStr, username)
	err := row.Scan(&u.User_name, &u.User_password, &u.Role)
	if err != nil {
		return c.String(http.StatusOK, "无此用户")
	}
	if password == u.User_password {
		token, err := middleware.CreateToken(u.User_name, u.Role == 1)
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		} else {
			return c.JSON(http.StatusOK, token)
		}
	} else {
		return c.String(http.StatusOK, "密码错误")
	}

}
