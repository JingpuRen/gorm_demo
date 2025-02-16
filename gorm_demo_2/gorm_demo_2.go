package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	// gorm默认将ID字段作为逐渐
	gorm.Model   // 这个里面就有id字段
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

// tip 如果我们想自己设置别的主键，我们可以通过tag来设置
type Animal struct {
	AnimalId int64 `gorm:"primary_key"` // 表示将这个字段设置为主键
	Name     string
	Age      int64
}

// tip 列名由字段名称进行下划线分割来生成
type Student struct {
	ID        uint      // column name is `id`
	Name      string    // column name is `name`
	Birthday  time.Time // column name is `birthday`
	CreatedAt time.Time // column name is `created_at`
}

// tip也可以使用结构体tag来指定列名
type Teacher struct {
	AnimalId int64     `gorm:"column:beast_id"`         // set column name to `beast_id`
	Birthday time.Time `gorm:"column:day_of_the_beast"` // set column name to `day_of_the_beast`
	Age      int64     `gorm:"column:age_of_the_beast"` // set column name to `age_of_the_beast`
}

func main() {
	// tip 修改默认的表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "JingpuRen_" + defaultTableName
	}

	db, err := gorm.Open("mysql", "root:666666@(localhost:3306)/db1?charset=utf8mb4&parseTime=True")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// 创建表，将表中的字段和代码结构体中的字段进行对应
	db.AutoMigrate(&User{})

	// 创建动物表
	db.AutoMigrate(&Animal{})

	// tip 如果不想自动创建结构体名称复数的表，我们可以通过指定创建的表的名称
	// db.Table("animal").CreateTable(&Animal{})
}
