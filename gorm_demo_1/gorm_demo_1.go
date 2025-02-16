package main

import "github.com/jinzhu/gorm"

type UserInfo struct {
	Id     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// tip 连接MySQL数据库
	db, err := gorm.Open("mysql", "root:666666@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=local")
	if err != nil {
		// tip 这段代码还有点不懂
		panic(err)
	}
	defer db.Close()

	// tip 创建表，将数据库表的字段和结构体中的字段进行一一对应
	db.AutoMigrate(&UserInfo{})

}
