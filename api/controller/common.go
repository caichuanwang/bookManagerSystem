package controller

import (
	"bookManagerSystem/untils"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var rdb *redis.Client

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
func ConnectRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6380",
		Password: "123456",
		DB:       0, // use default DB
		PoolSize: 20,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic("connection redis fail!" + err.Error())
	}

}
