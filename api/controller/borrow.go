package controller

import (
	"bookManagerSystem/modal"
	"bookManagerSystem/untils/sqlUntils"
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

var ctx = context.Background()

func CreateBorrow(c echo.Context) error {
	var u = new(modal.Borrow)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if ok, err := govalidator.ValidateStruct(u); !ok || err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//🔛事务
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var bookStock uint
	queryBookStockSql := "select bookStock from bookInfo where isbn = ?"
	stmt, err := db.Prepare(queryBookStockSql)
	if err != nil {
		_ = tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	row := stmt.QueryRow(u.Borrow_book_isbn)
	err = row.Scan(&bookStock)
	if err != nil {
		_ = tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if bookStock == 0 {
		_ = tx.Rollback()
		return c.JSON(http.StatusInternalServerError, modal.Err("no enough book"))
	}
	createSql := "insert into borrow(borrow_reader_id,borrow_book_isbn,is_borrow,borrow_time,agree_borrow_time,should_return_time,is_return,really_return_time) values(?,?,?,?,?,?,?,?)"
	stmt, err = db.Prepare(createSql)
	if err != nil {
		_ = tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()
	createRes, err := stmt.Exec(&u.Borrow_reader_id, &u.Borrow_book_isbn, 1, time.Now().Format("2006-01-02 15:04:05"), "", &u.Should_return_time, 1, "")
	if err != nil {
		_ = tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	createResRowsAffected, err := createRes.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	updateBookStock := "update bookInfo set bookStock = ? where isbn = ?"
	stmt, err = db.Prepare(updateBookStock)
	if err != nil {
		_ = tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	updateRes, err := stmt.Exec(bookStock-1, u.Borrow_book_isbn)
	if err != nil {
		_ = tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	updateResRowsAffected, err := updateRes.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if updateResRowsAffected != 1 || createResRowsAffected != 1 {
		_ = tx.Rollback()
		return c.JSON(http.StatusInternalServerError, modal.Err("insert or update sql error"))
	}
	//提交事务
	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	strCmd := rdb.Exists(ctx, u.Borrow_book_isbn)
	n := strCmd.Val()
	if n == 0 { //没有这个key
		if err := rdb.Set(ctx, u.Borrow_book_isbn, 0, 0).Err(); err != nil {
			return nil
		}
	} else { //自增1
		rdb.Incr(ctx, u.Borrow_book_isbn)
	}
	return c.JSON(http.StatusOK, modal.Success("ok"))
}

func QueryBorrowList(c echo.Context) error {
	var u = new(modal.QueryBorrowParams)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if ok, err := govalidator.ValidateStruct(u); !ok || err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var paramsMap = make(map[string]any)
	paramsMap["borrow_reader_name"] = u.Borrow_reader_name
	paramsMap["borrow_book_name"] = u.Borrow_book_name
	whereSql := sqlUntils.CreateWhereSql(paramsMap)
	orderSql := sqlUntils.CreateOrderSql(u.Order_by, u.Order_type)
	limitSql := sqlUntils.CreateLimitSql(u.Current, u.PageSize)
	querySql := fmt.Sprintf("select id,borrow_book_isbn,borrow_book_name,borrow_reader_id,borrow_reader_name,is_borrow,is_return,should_return_time,really_return_time,borrow_time,agree_borrow_time from borrow_with_name %s %s  %s", whereSql, orderSql, limitSql)
	stmt, err := db.Prepare(querySql)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var res []modal.BorrowWithName
	for rows.Next() {
		var a modal.BorrowWithName
		err = rows.Scan(&a.Id, &a.Borrow_book_isbn, &a.Borrow_book_name, &a.Borrow_reader_id, &a.Borrow_reader_name, &a.Is_borrow, &a.Is_return, &a.Should_return_time, &a.Really_return_time, &a.Borrow_time, &a.Agree_borrow_time)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		res = append(res, a)
	}
	queryCount := "select COUNT(1) from borrow_with_name"
	var a int
	db.QueryRow(queryCount).Scan(&a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.TableSucc(res, a))
}

func UpdateBorrowStatus(c echo.Context) error {
	updateSql := "update borrow set is_borrow=?,agree_borrow_time = ? where id = ?"
	stmt, err := db.Prepare(updateSql)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	_, err = stmt.Exec(modal.BORROW, time.Now().Format("2006-01-02 15:04:05"), c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("ok"))
}

func UpdateReturnStatus(c echo.Context) error {
	updateSql := "update borrow set is_return=?,really_return_time = ? where id = ?"
	stmt, err := db.Prepare(updateSql)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	_, err = stmt.Exec(modal.RETURN, time.Now().Format("2006-01-02 15:04:05"), c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("ok"))
}