package filter

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CustomFilter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//在这里处理拦截请求的逻辑
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return errors.New("Token isvalid")
		}
		token, _ := jwt.Parse(tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})
		fmt.Printf("%v  \n %v", token.Valid, token.Claims)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Printf("%v \n", claims["admin"])
			c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
			return next(c)
		} else {
			return c.String(http.StatusOK, "token is error")
			fmt.Println("token is error")
		}
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		//执行下一个中间件或者执行控制器函数, 然后返回执行结果
		return next(c)
	}
}
