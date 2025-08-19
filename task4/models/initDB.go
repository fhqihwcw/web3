package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:zhaoxj123@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields:                              true, //打印sql
		DisableForeignKeyConstraintWhenMigrating: true, //禁用外键约束
	})
	// DB.Debug()
	if err != nil {
		fmt.Println(err)
	}

	// 自动迁移
	err = DB.AutoMigrate(&User{}, &Post{}, &Comments{})
	if err != nil {
		fmt.Println("自动迁移失败:", err)
	}
}
