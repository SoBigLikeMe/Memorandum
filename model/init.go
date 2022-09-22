package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Database(connstring string) {
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		println("连接错误")
		println(err)
	} else {
		println("连接成功")
	}
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}

	//创建表时，表名末尾不加s
	db.SingularTable(true)

	//设置连接池
	db.DB().SetMaxIdleConns(20)

	//设置最大连接数
	db.DB().SetMaxOpenConns(100)

	//设置连接时间
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
}
