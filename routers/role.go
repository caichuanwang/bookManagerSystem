package routers

import (
	controller "bookManagementSystem/api/controller"
	"github.com/labstack/echo/v4"
)

func HandleRoleRoute(c *echo.Group) {
	go c.POST("/role/add", controller.CreateRole)
	go c.POST("/role/list", controller.QueryRole)
	go c.POST("/role/update", controller.UpdateRole)
	go c.GET("/role/delete", controller.DeleteRole)
}
