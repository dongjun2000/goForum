package models

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

// DB 变量代表数据库连接池
var Db *sql.DB

// 数据库连接初始化方法，在Web应用启动时自动初始化数据库连接
func init() {
	var err error

	// 配置数据源信息，用于定义如何连接数据库
	config := mysql.Config{
		User: "root",
		Passwd: "",
		Addr: ":3306",
		Net: "tcp",
		DBName: "goforum",
		AllowNativePasswords: true,
	}

	Db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	// 设置最大连接数
	Db.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	Db.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	Db.SetConnMaxLifetime(5 * time.Minute)
	return
}




