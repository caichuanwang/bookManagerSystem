package main

import (
	"bookManagerSystem/api/controller"
	_ "bookManagerSystem/docs"
	"bookManagerSystem/routers"
	"bookManagerSystem/untils"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @host petstore.swagger.io
// @BasePath /v1
func main() {
	e := echo.New()
	e.Static("/", "static")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Any("/*", echoSwagger.WrapHandler)
	//自定义中间件
	//e.Use(filter.CustomFilter)
	e.POST("/login", controller.HandleLoginController)
	//初始化连接数据库实例
	go controller.DriverMySQL()

	r := e.Group("/v1")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS512", SigningKey: []byte("secret")}))
	routers.InitRouter(r)

	//初始化路由
	//开启服务
	port := untils.ReadCon("base", "port")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
