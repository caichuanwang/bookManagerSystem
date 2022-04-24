package routers

import (
	"bookManagerSystem/api/controller"
	"github.com/labstack/echo/v4"
)

func HandleLoginRoute(c *echo.Echo) {
	c.POST("/login", controller.HandleLoginController)
}
