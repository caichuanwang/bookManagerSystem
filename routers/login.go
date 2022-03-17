package routers

import (
	controller "bookManagementSystem/api/controller"
	"github.com/labstack/echo/v4"
)

func HandleLoginRoute(c *echo.Echo) {
	c.POST("/login", controller.HandleLoginController)
}
