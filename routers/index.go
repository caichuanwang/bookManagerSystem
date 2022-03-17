package routers

import (
	"github.com/labstack/echo/v4"
)

func InitRouter(c *echo.Group) {
	//HandleLoginRoute(c)
	HandleUserRoute(c)
	HandleRoleRoute(c)
}
