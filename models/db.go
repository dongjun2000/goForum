package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
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
		ParseTime: true,
		AllowNativePasswords: true,
	}

	fmt.Println(config.FormatDSN())
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

	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	return
}

// create a random UUID with from RFC 4122
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatal("Cannot generate UUID", err)
	}

	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// sha1哈希
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}



