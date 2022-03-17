package main

import (
	controller "bookManagementSystem/api/controller"
	customMiddleware "bookManagementSystem/api/middleware"
	"bookManagementSystem/filter"
	"bookManagementSystem/routers"
	"bookManagementSystem/untils"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.Static("/", "static")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//自定义中间件
	e.Use(filter.CustomFilter)

	e.POST("/login", controller.HandleLoginController)
	e.GET("/", accessible)
	//初始化连接数据库实例
	go controller.DriverMySQL()
	config := middleware.JWTConfig{
		Claims:     &customMiddleware.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	r := e.Group("/v1", filter.CustomFilter)

	r.Use(middleware.JWTWithConfig(config))
	r.GET("", restricted)
	routers.InitRouter(r)

	//初始化路由
	//开启服务
	port := untils.ReadCon("base", "port")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func restricted(c echo.Context) error {
	fmt.Println("211111")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*customMiddleware.JwtCustomClaims)
	//log.InfoDump(claims, "claims")
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
func accessible(c echo.Context) error {
	fmt.Println("4444444")
	return c.String(http.StatusOK, "Accessible")
}
func login(c echo.Context) error {
	username := c.FormValue("user_name")
	password := c.FormValue("password")
	fmt.Println(username, password)
	if username == "admin" && password == "admin" {

		// Set custom claims
		claims := &customMiddleware.JwtCustomClaims{
			"admin",
			true,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}
