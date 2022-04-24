package controller

import (
	"bookManagerSystem/api/middleware"
	"bookManagerSystem/modal"
	"bookManagerSystem/untils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"net/http"
)

var database = untils.ReadCon("mysql", "databaseName")

type LoginParams struct {
	User_name     string `json:"user_name"`
	User_password string `json:"user_password"`
}
type LoginReturn struct {
	Token    any    `json:"token"`
	userName string `json:"userName"`
}

func HandleLoginController(c echo.Context) error {
	var req = new(LoginParams)
	c.Bind(req)

	//var u = modal.NewUser()
	//defer c.Request().Body.Close()
	//b, err := ioutil.ReadAll(c.Request().Body)
	//if err != nil {
	//	log.Println("failed reading the request body")
	//	return c.String(http.StatusInternalServerError, err.Error())
	//}
	//json.Unmarshal(b, &u)
	//上面注释的方式是使用流读取参数也是可以的
	var u1 = modal.NewUser()
	queryStr := fmt.Sprintf("select user_name,user_password,role from %s where user_name = ?", "user")
	row := db.QueryRow(queryStr, req.User_name)
	err := row.Scan(&u1.User_name, &u1.User_password, &u1.Role)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, modal.Err("user is not exist"))
	}
	if req.User_password == u1.User_password {
		token, err := middleware.CreateToken(u1.User_name, u1.Role == "1")
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		} else {
			return c.JSON(http.StatusOK, modal.Success(&LoginReturn{
				Token:    token["token"],
				userName: u1.User_name,
			}))
		}
	} else {
		return c.JSON(http.StatusUnauthorized, modal.Err("password error"))
	}

}
