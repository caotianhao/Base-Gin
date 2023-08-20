package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	//这个是驱动，必须导入，且前面加下划线
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type userInfo struct {
	//这里最终建成的表名是 user_infos，是单词拆开加下划线分割然后最后加复数
	//后续会学到自定义表名
	//往后这里面都得大写，再也不是仅本地读取了！！！
	Id   uint
	Name string
	Age  int
}

// 什么是 ORM
// Object		对象：程序中的对象/实例，例如 go 中的结构体实例
// Relational	关系：关系数据库，例如MySQL
// Mapping		映射
// 优点：提高开发效率
// 缺点：牺牲执行性能，牺牲灵活性，弱化 SQL 能力
func main() {
	//连接 MySQL 数据库              用户名:密码  @(ip及端口)       /数据库名?字符集       &解析时间类型    &本地时间
	db, err := gorm.Open("mysql", "root:824416@(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败，err =", err)
		return
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("关闭失败，err =", err)
			return
		}
	}(db)

	//创建表，自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&userInfo{})

	//创建数据行
	u1 := userInfo{1, "alice", 20}
	u2 := userInfo{2, "bob", 21}
	u3 := userInfo{3, "cindy", 11}
	u4 := userInfo{4, "dad", 18}
	//传结构体也能传，但是太大了，所以传指针
	//db.Create(u1)
	db.Create(&u1)
	db.Create(&u2)
	db.Create(&u3)
	db.Create(&u4)

	//创建查询
	var u userInfo
	//查第一条数据保存到 u 里
	db.First(&u)
	fmt.Println(u) //{1 alice 20}

	//更新
	db.Model(&u3).Update("Age", 3)

	//删除
	db.Delete(&u4)
	//+----+-------+------+
	//| id | name  | age  |
	//+----+-------+------+
	//|  1 | alice |   20 |
	//|  2 | bob   |   21 |
	//|  3 | cindy |    3 |
	//+----+-------+------+
}
