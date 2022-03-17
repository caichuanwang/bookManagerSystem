package routers

import (
	controller2 "bookManagementSystem/api/controller"
	"github.com/labstack/echo/v4"
)

func HandleUserRoute(c *echo.Group) {
	go c.POST("/user/add", controller2.CreateAddUser)
	go c.POST("/user/list", controller2.QueryUser)
	go c.POST("/user/update", controller2.UpdateUser)
	go c.GET("/user/delete", controller2.DeleteUser)
}
