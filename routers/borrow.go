package routers

import (
	"bookManagerSystem/api/controller"
	"github.com/labstack/echo/v4"
)

func HandleBorrowRouter(e *echo.Group) {
	e.POST("/borrow/add", controller.CreateBorrow)
	e.POST("/borrow/list", controller.QueryBorrowList)
	e.GET("/borrow/borrow/status", controller.UpdateBorrowStatus)
	e.GET("/borrow/return/status", controller.UpdateReturnStatus)
	e.GET("/borrow/top", controller.GetTopBookList)
}
