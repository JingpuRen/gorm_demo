package main

import (
	"fmt"
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
	// tip 时区问题已经解决，注意是loc=Local，后面的Local是要大写的！！！
	db, err := gorm.Open("mysql", "root:666666@(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		// tip 这段代码还有点不懂
		panic(err)
	}
	defer db.Close()

	// tip 创建表，将数据库表的字段和结构体中的字段进行一一对应
	//db.AutoMigrate(&UserInfo{})

	// tip 创建数据行
	//u1 := UserInfo{1, "JingpuRen", "男", "coding"}
	//db.Create(&u1)

	// tip 查询数据行
	var u2 UserInfo
	db.First(&u2)
	fmt.Println(u2.Name)

	// tip 更新数据行中的某个字段
	db.Model(&u2).Update("hobby", "跑步")

	// tip 删除数据库中的某行记录
	db.Delete(&u2)
}
