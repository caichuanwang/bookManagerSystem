package routers

import (
	"bookManagerSystem/api/controller"
	"github.com/labstack/echo/v4"
)

func HandleBookListRouter(e *echo.Group) {
	e.POST("/bookList/add", controller.CreateBookList)
	e.GET("/bookList/options", controller.QueryBookListOptions)
	e.POST("/bookList/set2BookList", controller.SetBook2BookList)
	e.POST("/bookList/list", controller.QueryBookListList)
	e.DELETE("/bookList/delete", controller.DeleteBookListList)
}
