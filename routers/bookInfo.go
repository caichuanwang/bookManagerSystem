package routers

import (
	"bookManagerSystem/api/controller"
	"github.com/labstack/echo/v4"
)

func HandleBookInfoRouter(e *echo.Group) {
	e.POST("/bookInfo/add", controller.CreateBook)
	e.POST("/bookInfo/update", controller.UpdateBookInfo)
	e.POST("/bookInfo/list", controller.QueryBookList)
	e.DELETE("/bookInfo/delete", controller.DeleteBookInfo)
}
