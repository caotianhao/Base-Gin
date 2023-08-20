package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User :ID 为主键
// 这样的话通常创建一条记录时需要指定 Name 和 Age 字段
// 但假设不指定 Name 字段，希望它有一个默认值，可以使用 tag
type User struct {
	ID   int64
	Name string `gorm:"default:'no_name'"` //若不设默认值，则默认值为空，不是 null
	Age  int64  //默认值是0
	//通过 tag 设定的字段默认值，在创建记录的时候的 sql 语句会排除没有值或为零值的字段
	//在将记录插到数据库中之后，gorm 会从数据库加载那些字段的默认值
	//u := User{Age: 222}           对应的 name 字段为 no_name
	//u := User{Name: "", Age: 222} 对应的 name 字段仍为 no_name
	//如果想避免这种情况，可以使用指针
}

// User2 使用指针避免把空串改变设为默认值
type User2 struct {
	ID   int64
	Name *string `gorm:"default:'no_name'"` //若不设默认值，则默认值为空，不是 null
	Age  int64   //默认值是0
	//在初始化时
	//u := User{Name: new(string), Age: 11}
}

// User3 使用 Scanner/Valuer 避免把空串改变设为默认值
type User3 struct {
	ID   int64
	Name sql.NullString `gorm:"default:'no_name'"` //若不设默认值，则默认值为空，不是 null
	Age  int64          //默认值是0
	//在初始化时
	//u := User{Name: sql.NullString{"", true}, Age: 11}
}

func main() {
	//gorm 是没有办法为我们创建数据库的，我们必须先手动 create database gorm21;
	db, err := gorm.Open("mysql", "root:824416@(localhost:3306)/gorm21?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("gorm.Open err =", err)
		return
	}

	//该语句可以是结构体名字变为表名时不变复数
	db.SingularTable(true)

	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("db.Close err =", err)
			return
		}
	}(db) 

	db.AutoMigrate(&User{})

	//ID 为主键，可以不用写
	u := User{Name: "cth", Age: 925}
	//判断该主键是否为空
	fmt.Println(db.NewRecord(&u))
	//可以使用 Debug() 查看建表语句
	//db.Debug().Create(&u)
	db.Create(&u)
	fmt.Println(db.NewRecord(&u))
	//u2 := User2{Age: 925}
	//u3 := User3{Age: 925}
	//fmt.Println(u2, u3)
}
