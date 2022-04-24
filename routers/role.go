package routers

import (
	"bookManagerSystem/api/controller"
	"github.com/labstack/echo/v4"
)

func HandleRoleRoute(c *echo.Group) {
	c.POST("/role/list", controller.QueryRole)
	c.GET("/role/option", controller.QueryRoleOptions)
	c.POST("/role/add", controller.CreateRole)
	c.POST("/role/update", controller.UpdateRole)
	c.DELETE("/role/delete", controller.DeleteRole)

}
