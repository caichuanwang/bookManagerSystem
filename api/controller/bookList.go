package controller

import (
	"bookManagerSystem/modal"
	"bookManagerSystem/untils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

func CreateBookList(c echo.Context) error {
	claims, err := untils.EncodingUser(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var u modal.BookList
	var user modal.User
	gdb.Where("id = ?", claims.Id).First(&user)
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = gdb.Create(&u).Association("User").Append(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("ok"))
}

func QueryBookListOptions(c echo.Context) error {
	claims, err := untils.EncodingUser(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var a []modal.BookList
	gdb.Select("id", "name", "user_id").Where("user_id = ?", claims.Id).Find(&a)
	var options []modal.SelectOption
	for _, v := range a {
		options = append(options, modal.SelectOption{
			Label: v.Name,
			Value: v.ID,
		})
	}
	return c.JSON(http.StatusOK, modal.Success(options))
}

func SetBook2BookList(c echo.Context) error {
	var u modal.SetBook2BookListParams
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var a modal.BookInfo
	gdb.Where("isbn = ?", u.Isbn).First(&a)
	var b []modal.BookList
	gdb.Where("id IN ? ", u.BookLists).Find(&b)
	err := gdb.Model(&a).Association("BookList").Replace(&b)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var d []modal.BookList
	gdb.Model(&a).Preload("BookInfo", "isbn = ?", a.Isbn).Association("BookList").Find(&d)
	return c.JSON(http.StatusOK, modal.Success("ok"))
}

func IsCollect(c echo.Context) error {
	claims, err := untils.EncodingUser(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var u modal.User
	var b []modal.BookList
	gdb.Model(&u).Where("id = ?", claims.Id).First(&u)
	err = gdb.Model(&u).Association("BookList").Find(&b)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}
	if len(b) == 0 {
		return c.JSON(http.StatusOK, modal.Success(false))
	}
	var book []modal.BookInfo
	err = gdb.Model(&b).Association("BookInfo").Find(&book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	for _, info := range book {
		if info.Isbn == c.QueryParam("isbn") {
			return c.JSON(http.StatusOK, modal.Success(true))
		}
	}
	return c.JSON(http.StatusOK, modal.Success(false))
}

func QueryBookListList(c echo.Context) error {
	var p modal.QueryBookListParams
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var bookLists []modal.BookList
	gdb.Model(&bookLists).Where("name LIKE ?", "%"+p.Name+"%").Preload("BookInfo").Limit(p.PageSize).Offset((p.Current - 1) * p.PageSize).Find(&bookLists)
	var t int64
	gdb.Model(&bookLists).Count(&t)
	return c.JSON(http.StatusOK, modal.TableSucc(bookLists, int(t)))
}

func DeleteBookListList(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var bl = modal.BookList{ID: uint(id)}
	gdb.Delete(&bl)
	err = gdb.Model(&bl).Association("BookInfo").Clear()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	err = gdb.Model(&bl).Association("User").Clear()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	log.Fatal()
	return c.JSON(http.StatusOK, modal.Success("ok"))
}
