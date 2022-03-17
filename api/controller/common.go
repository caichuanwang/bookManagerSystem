package controller

import (
	"bookManagementSystem/untils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DriverMySQL() {
	driverName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", untils.ReadCon("mysql", "rootName"), untils.ReadCon("mysql", "password"), untils.ReadCon("mysql", "location"), untils.ReadCon("mysql", "port"), untils.ReadCon("mysql", "databaseName"))
	sqlSession, err := sql.Open("mysql", driverName)
	if err != nil {
		panic("open database error!")
	}
	//犯了大错了 这里关闭了连接  后面怎么使用db呢？？？？？？？
	//defer sqlSession.Close()
	err = sqlSession.Ping()
	if err != nil {
		panic("connection database fail!")
	}
	db = sqlSession

}
