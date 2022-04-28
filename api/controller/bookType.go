package controller

import (
	"bookManagerSystem/modal"
	"bookManagerSystem/untils/sqlUntils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateBookType(c echo.Context) error {
	var u = new(modal.BookType)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if ok, err := govalidator.ValidateStruct(u); err != nil || !ok {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	sqlString := "insert into  book_type (typeName,path,level,branch,remake) values (?,?,?,?,?)"
	stmt, err := db.Prepare(sqlString)
	defer stmt.Close()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	_, err = stmt.Exec(&u.TypeName, &u.Path, &u.Level, &u.Branch, &u.Remake)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("ok"))
}

func GetBookTypeList(c echo.Context) error {
	var u = new(modal.BookTypeReqParams)
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	orderBySql := sqlUntils.CreateOrderSql(u.Order_by, u.Order_type)
	limitSql := sqlUntils.CreateLimitSql(u.Current, u.PageSize)
	querySql := fmt.Sprintf("select id,typeName,path,level,branch,remake from book_type %s %s", orderBySql, limitSql)
	stmt, err := db.Prepare(querySql)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	var res []modal.BookType
	for rows.Next() {
		var tempBook modal.BookType
		err := rows.Scan(&tempBook.Id, &tempBook.TypeName, &tempBook.Path, &tempBook.Level, &tempBook.Branch, &tempBook.Remake)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		res = append(res, tempBook)
	}
	queryCount := "select count(id) from book_type"
	var a int
	db.QueryRow(queryCount).Scan(&a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.TableSucc(res, a))
}

func UpdateBookType(c echo.Context) error {
	var b = new(modal.BookType)
	if err := c.Bind(&b); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	updateSql := "update book_type set typeName= ?,path=?,level=?,branch=?,remake=? where id = ?"
	stmt, err := db.Prepare(updateSql)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	_, err = stmt.Exec(&b.TypeName, &b.Path, &b.Level, &b.Branch, &b.Remake, &b.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("ok"))
}

func DeleteBookType(c echo.Context) error {
	deleteSql := "delete from book_type where id = ?"
	stmt, err := db.Prepare(deleteSql)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	_, err = stmt.Exec(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("ok"))
}
