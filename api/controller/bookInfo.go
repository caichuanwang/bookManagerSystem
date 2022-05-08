package controller

import (
	"bookManagerSystem/modal"
	"bookManagerSystem/untils"
	"bookManagerSystem/untils/sqlUntils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"path"
)

func CreateBook(c echo.Context) error {
	var b modal.CreateBookInfoParams
	if err := c.Bind(&b); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if ok, err := govalidator.ValidateStruct(&b); !ok && err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var imgPath string
	if b.Photo != nil {
		//存图片到文件中
		var imgStr = b.Photo[0]["thumbUrl"].(string)
		var uuidStr = uuid.New().String()
		imgPath = untils.ReadCon("book", "imgPath") + "/" + uuidStr + path.Ext(b.Photo[0]["name"].(string))
		_ = untils.Base642Img(imgStr[22:], untils.ReadCon("book", "imgPath"), uuidStr+path.Ext(b.Photo[0]["name"].(string)))
	}
	createSql := "insert into bookInfo(isbn,bookName,author,translator,publisher,publishTime,bookStock,price,typeId,context,photo,pageNum)values(?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(createSql)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	_, err = stmt.Exec(b.Isbn, b.BookName, b.Author, b.Translator, b.Publisher, b.PublishTime, b.BookStock, b.Price, b.TypeId, b.Context, imgPath, b.PageNum)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("ok"))
}

func QueryBookList(c echo.Context) error {
	var u = new(modal.QueryBookInfoParams)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	var paramMap = make(map[string]any)
	paramMap["bookName"] = u.BookName
	whereCon := sqlUntils.CreateWhereSql(paramMap)
	orderBySql := sqlUntils.CreateOrderSql(u.Order_by, u.Order_type)
	LimitSql := sqlUntils.CreateLimitSql(u.Current, u.PageSize)
	queryStr := fmt.Sprintf("select isbn,bookName,author,translator,publisher,publishTime,bookStock,price,typeId,context,photo,pageNum,(select typeName from book_type where id = bookInfo.typeId ) as typeName from bookInfo %s %s %s ", whereCon, orderBySql, LimitSql)
	fmt.Println(queryStr)
	stmt, err := db.Prepare(queryStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var res []modal.BookInfo
	for rows.Next() {
		var a modal.BookInfo
		err = rows.Scan(&a.Isbn, &a.BookName, &a.Author, &a.Translator, &a.Publisher, &a.PublishTime, &a.BookStock, &a.Price, &a.TypeId, &a.Context, &a.Photo, &a.PageNum, &a.TypeName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		res = append(res, a)
	}
	queryCount := "select COUNT(1) from bookInfo"
	var a int
	db.QueryRow(queryCount).Scan(&a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.TableSucc(res, a))
}

func UpdateBookInfo(c echo.Context) error {
	var u = new(modal.BookInfo)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	ok, err := govalidator.ValidateStruct(u)
	if err != nil || !ok {
		return c.String(http.StatusBadRequest, err.Error())
	}
	updateUserSQL := "update bookInfo set isbn = ?,bookName = ? ,author = ? ,translator = ?,publisher=?,publishTime = ?,bookStock=?,price=?,typeId=?,context=?,pageNum=?,photo=? where isbn = ?"
	stmt, err := db.Prepare(updateUserSQL)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	_, err = stmt.Exec(&u.Isbn, &u.BookName, &u.Author, &u.Translator, &u.Publisher, &u.PublishTime, &u.BookStock, &u.Price, &u.TypeId, &u.Context, &u.PageNum, &u.Photo, &u.Isbn)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("ok"))
}

func DeleteBookInfo(c echo.Context) error {
	deleteUserSql := "delete from bookInfo where isbn = ?"
	stmt, err := db.Prepare(deleteUserSql)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	_, err = stmt.Exec(c.QueryParam("isbn"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("ok"))
}
