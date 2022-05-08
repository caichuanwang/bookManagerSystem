package routers

import (
	"bookManagerSystem/api/controller"
	"github.com/labstack/echo/v4"
)

func HandleBookTypeRouter(e *echo.Group) {
	e.POST("/bookType/add", controller.CreateBookType)
	e.POST("/bookType/list", controller.GetBookTypeList)
	e.POST("/bookType/update", controller.UpdateBookType)
	e.DELETE("/bookType/delete", controller.DeleteBookType)
	e.GET("/bookType/treeList", controller.GetBookTypeWithTree)
}
