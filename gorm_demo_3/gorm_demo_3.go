package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Person struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	db, err := gorm.Open("mysql", "root:666666@(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	// tip 第一次运行时需要使用这段代码 创建表
	// tip 第一次代码之后也不用注释掉，第一次负责创建表，之后就负责将数据库表的结构和对应的结构体的结构对应起来了
	db.Table("new_person").AutoMigrate(&Person{})

	// 插入记录
	u1 := Person{Name: "sing", Age: 18}

	// tip 这里一定要注意，如果你之前创建表使用Table指定了表名，那么之后我们在创建数据的时候也是要指定表名的！！
	// tip 这里真的是很容易错的地方！！！
	db.Table("new_person").Create(&u1)

	u2 := Person{Name: "haha", Age: 20}
	// tip 此处同理，这里是非常容易错的地方！！！一定要记住要加上Table！！！
	db.Table("new_person").Create(&u2)

	defer db.Close()

	// 查询
}
