package routers

import (
	"bookManagerSystem/api/controller"
	"github.com/labstack/echo/v4"
)

func HandleUserRoute(c *echo.Group) {
	c.POST("/user/list", controller.QueryUser)
	c.POST("/user/add", controller.CreateAddUser)
	c.POST("/user/update", controller.UpdateUser)
	c.DELETE("/user/delete", controller.DeleteUser)

}
