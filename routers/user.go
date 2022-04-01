package routers

import (
	controller "bookManagementSystem/api/controller"
	"github.com/labstack/echo/v4"
)

func HandleUserRoute(c *echo.Group) {
	go c.POST("/user/add", controller.CreateAddUser)
	go c.POST("/user/list", controller.QueryUser)
	go c.POST("/user/update", controller.UpdateUser)
	go c.DELETE("/user/delete", controller.DeleteUser)
}
