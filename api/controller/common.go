package controller

import (
	"bookManagerSystem/modal"
	"bookManagerSystem/untils"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jordan-wright/email"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"net/smtp"
	"time"
)

var db *sql.DB
var rdb *redis.Client
var emailPool *email.Pool
var gdb *gorm.DB

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
		panic(fmt.Sprintf("connection database fail!:%s", err.Error()))
	}
	db = sqlSession

}
func ConnectRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     untils.ReadCon("redis", "addr"),
		Password: untils.ReadCon("redis", "password"),
		DB:       0, // use default DB
		PoolSize: 20,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic("connection redis fail!" + err.Error())
	}

}

func InitEMailPool() {
	pool, err := email.NewPool("smtp.qq.com:25", 10, smtp.PlainAuth("", untils.ReadCon("qqEmail", "username"), untils.ReadCon("qqEmail", "password"), "smtp.qq.com"))
	if err != nil {
		panic("init email error :" + err.Error())
	}
	emailPool = pool
}

func TimerSendEmail() {
	c := cron.New()
	c.AddFunc("@daily", func() {
		now := time.Now().Format("2006-01-02")
		querySql := "select borrow_book_name,borrow_reader_name,email from borrow_with_name where  should_return_time = ?"
		stmt, err := db.Prepare(querySql)
		if err != nil {
			log.Fatal("get email error" + err.Error())
		}
		rows, err := stmt.Query(now)
		if err != nil {
			log.Fatal("get email error" + err.Error())
		}
		var bookName string
		var userName string
		var userEmail string
		for rows.Next() {
			err = rows.Scan(&bookName, &userName, &userEmail)
			if err != nil {
				log.Fatal("get email error" + err.Error())
			}
			text := fmt.Sprintf("您借的%s期限与今日到期，请经快归还书籍", bookName)
			go SendEmail([]string{userEmail}, "还书提醒", text)
		}
	})
	c.Start()
}

func ConnectMysqlWithGorm() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		untils.ReadCon("mysql", "rootName"),
		untils.ReadCon("mysql", "password"),
		untils.ReadCon("mysql", "location"),
		untils.ReadCon("mysql", "port"),
		untils.ReadCon("mysql", "databaseName"))
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 191,
	}), &gorm.Config{
		SkipDefaultTransaction: false, //关闭默认事务，gorm默认每次crud都是在事务中完成的
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "g_", //表名的前缀
			SingularTable: true, //表名是否使用单数，还是复数
			NoLowerCase:   false,
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 逻辑外键，（提倡代码里自动外键关联）

	})
	if err != nil {
		panic("gorm connect mysql error" + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("gorm connect mysql error" + err.Error())
	}
	sqlDB.SetMaxOpenConns(100)                                                //设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxIdleTime(10)                                              //设置空闲连接池中连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour)                                       //设置了连接可复用的最大时间。
	err = db.AutoMigrate(&modal.BookList{}, &modal.BookInfo{}, &modal.User{}) //自动创建表
	if err != nil {
		log.Fatal(err.Error())
	}
	gdb = db
}
