package controller

import (
	"bookManagerSystem/modal"
	"bookManagerSystem/untils/sqlUntils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// CreateBookType @Summary 新增📖类型
// @Description 新增📖类型
// @Accept json
// @Param user body  modal.BookType true "新增📖的数据"
// @Success 200 {object} modal.Result
// @Router /v1/bookType/add [post]
func CreateBookType(c echo.Context) error {
	var u = new(modal.BookType)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if ok, err := govalidator.ValidateStruct(u); err != nil || !ok {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	sqlString := "insert into  book_type (typeName,level,pId,remake) values (?,?,?,?)"
	stmt, err := db.Prepare(sqlString)
	defer stmt.Close()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	_, err = stmt.Exec(&u.TypeName, &u.Level, &u.PId, &u.Remake)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("ok"))
}

// GetBookTypeList @Summary 获取图书分类列表
// @Description 获取图书分类列表
// @Accept json
// @Param user body  modal.BookTypeReqParams true "查询参数"
// @Success 200 {object} modal.TableResult
// @Router /v1/bookType/list [post]
func GetBookTypeList(c echo.Context) error {
	var u = new(modal.BookTypeReqParams)
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var paramMap = make(map[string]any)
	paramMap["typeName"] = u.TypeName
	paramMap["level"] = u.Level
	paramMap["pName"] = u.PName
	whereCon := sqlUntils.CreateWhereSql(paramMap)
	fmt.Println(whereCon)
	orderBySql := sqlUntils.CreateOrderSql(u.Order_by, u.Order_type)
	limitSql := sqlUntils.CreateLimitSql(u.Current, u.PageSize)
	querySql := fmt.Sprintf("select id,typeName,level ,remake , pName, pId  from book_type_with_pName %s %s %s", whereCon, orderBySql, limitSql)
	stmt, err := db.Prepare(querySql)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	var res []modal.BookTypeResult
	for rows.Next() {
		var tempBook modal.BookTypeResult
		err := rows.Scan(&tempBook.Id, &tempBook.TypeName, &tempBook.Level, &tempBook.Remake, &tempBook.PName, &tempBook.PId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		res = append(res, tempBook)
	}
	queryCount := fmt.Sprintf("select count(id) from book_type_with_pName %s", whereCon)
	var a int
	err = db.QueryRow(queryCount).Scan(&a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.TableSucc(res, a))
}

// UpdateBookType
// @Description 更新图书分类列表
// @Accept json
// @Param user body  modal.BookType true "更新参数"
// @Success 200 {object} modal.Result
// @Router /v1/bookType/update [post]
func UpdateBookType(c echo.Context) error {
	var b = new(modal.BookType)
	if err := c.Bind(&b); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	updateSql := "update book_type set typeName= ?,level=?,pId=?,remake=? where id = ?"

	stmt, err := db.Prepare(updateSql)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	_, err = stmt.Exec(&b.TypeName, &b.Level, &b.PId, &b.Remake, &b.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modal.Success("ok"))
}

// DeleteBookType
// @Description 删除图书分类
// @Accept json
// @Param user body  string true "删除图书分类的🆔"
// @Success 200 {object} modal.Result
// @Router /v1/bookType/delete [delete]
func DeleteBookType(c echo.Context) error {
	querySql := "select count(1) from book_type where pId = ?"
	s, err := db.Prepare(querySql)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	res := s.QueryRow(c.QueryParam("id"))
	var a uint
	err = res.Scan(&a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if a != 0 {
		return c.JSON(http.StatusInternalServerError, modal.Err("this type has children type,can not delete"))
	}
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

// GetBookTypeWithTree
// @Description 返回树形结构的图书类型
// @Accept json
// @Success 200 {object} []modal.TreeBookType
// @Router /v1/bookType/treeList [get]
func GetBookTypeWithTree(c echo.Context) error {
	querySql := "select id,typeName,pId,remake,level from book_type"
	stmt, err := db.Prepare(querySql)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	var res []modal.ItemBookType

	for rows.Next() {
		var b = new(modal.ItemBookType)
		err = rows.Scan(&b.Id, &b.TypeName, &b.PId, &b.Remake, &b.Level)
		res = append(res, *b)
	}
	rows.Close()
	tree := Array2Tree(res)
	return c.JSON(http.StatusOK, tree)
}

type treeMapTree map[string]*modal.TreeBookType

func Array2Tree(data []modal.ItemBookType) []*modal.TreeBookType {
	var tree []*modal.TreeBookType
	var treeMap = make(treeMapTree)
	for _, item := range data {
		var temp modal.TreeBookType
		temp.Id = item.Id
		temp.PId = item.PId
		temp.TypeName = item.TypeName
		temp.Remake = item.Remake
		temp.Level = item.Level
		temp.Children = nil
		if item.PId == 0 {
			tree = append(tree, &temp)
		} else {
			treeMap[strconv.Itoa(int(temp.PId))].Children = append(treeMap[strconv.Itoa(int(temp.PId))].Children, &temp)
		}
		treeMap[strconv.Itoa(int(temp.Id))] = &temp
	}
	return tree
}
