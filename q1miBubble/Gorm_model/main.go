package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//在使用 ORM 工具时，通常我们需要在代码中定义模型与数据库中的数据表进行映射
//在 GORM 中模型通常是正常定义的结构体、基本的 go 类型或它们的指针
//同时也支持 sql.Scanner 及 driver.Valuer 接口
//为了方便模型定义，gorm 内置了一个 gorm.Model 结构体
//是一个包含了 ID，CreatedAt，UpdatedAt，DeletedAt 四个字段的结构体
//其中 CreatedAt 字段的值为初次创建记录的时间，可以使用 Update 更改
//UpdatedAt 则是每次更新记录的时间
//DeletedAt 是指调用 delete 删除记录时，值为当前时间

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(120);unique_index"`
	Role         string  `gorm:"size:255"`        //设置字段大小为 255
	MemberNumber *string `gorm:"unique;not null"` //设置会员号唯一且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  //设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      //给 address 字段创建名为 addr 的索引
	IgnoreMe     int     `gorm:"-"`               //忽略本字段
}

//gorm 会默认使用名为 id 的字段作为主键
//加如下的 tag 可以指定主键

type Animal struct {
	AnimalID int `gorm:"primary_key"`
	Name     string
	Age      int `gorm:"column:animal_age"` //自己指定列名
}

//表名默认就是结构体名字的复数
//表名是可以使用 TableName() 自己设置的，以上面的 animal 为例
//我想要表名为 animalChange
//这样执行之后，会创建一个和 animals 一样的新表，并不会把旧的 animals 改为 animalChange

//func (Animal) TableName() string {
//	return "animalChange"
//}

//也可以根据结构体具体的值，给予不同的表名

func (u User) TableName() string {
	if u.Name == "admin" {
		return "admin_users"
	} else {
		return "normal_users"
	}
}

func main() {
	//连接 mysql
	db, err := gorm.Open("mysql", "root:824416@(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("sql open failed, err =", err)
		return
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("close mysql failed, err =", err)
			return
		}
	}(db)

	//gorm 还可以修改默认的表名规则
	//比如我某个项目的表都需要以某个前缀开始
	//但这个只会修改默认的表名规则，如果通过 TableName 指定的则不受影响
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "myPrefix_" + defaultTableName
	}

	//禁用表名自动复数（加 s）
	db.SingularTable(true)

	//自动迁移
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

	//也可以使用 Table() 直接指定名字
	//使用 my_user 作为表名，创建对应于 User 结构体的表
	db.Table("my_user").CreateTable(&User{})
}
