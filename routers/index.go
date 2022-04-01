package routers

import (
	"github.com/labstack/echo/v4"
)

func InitRouter(c *echo.Group) {
	HandleUserRoute(c)
	HandleRoleRoute(c)
}
