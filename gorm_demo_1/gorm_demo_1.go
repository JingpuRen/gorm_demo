package main

import (
	"github.com/jinzhu/gorm"
	// tip 一定要记得引入对应的数据库驱动啊！！！
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	Id     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// tip 连接MySQL数据库
	// todo 这个时区设立的还有问题，后续还待解决
	db, err := gorm.Open("mysql", "root:666666@(localhost:3306)/db1?charset=utf8mb4&parseTime=True")
	if err != nil {
		// tip 这段代码还有点不懂
		panic(err)
	}
	defer db.Close()

	// tip 创建表，将数据库表的字段和结构体中的字段进行一一对应
	db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{1, "JingpuRen", "男", "coding"}
	db.Create(&u1)

}
