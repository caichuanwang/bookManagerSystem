package routers

import (
	"bookManagerSystem/api/controller"
	"github.com/labstack/echo/v4"
)

func HandleBookTypeRouter(e *echo.Group) {
	e.POST("/bookType/add", controller.CreateBookType)
	e.POST("/bookType/list", controller.GetBookTypeList)
	e.POST("/bookType/update", controller.UpdateBookType)
	e.GET("/bookType/delete", controller.DeleteBookType)
}
