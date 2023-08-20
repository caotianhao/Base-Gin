package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

//待办事项结构体

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	//创建并连接数据库
	dsn := "root:824416@(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Open mysql failed err =", err)
		return
	}

	//关闭数据库
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("defer close db failed err =", err)
			return
		}
	}(db)

	db.AutoMigrate(&Todo{})

	r := gin.Default()

	//告诉 gin 模板文件引用的静态文件去哪里找
	r.Static("/static", "static")

	//告诉 gin 去哪里找 index.html
	//* 表示任何文件
	r.LoadHTMLGlob("template/*")

	r.GET("/bubble", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		//添加待办事项
		v1Group.POST("/todo", func(c *gin.Context) {
			//前端页面填写待办事项 点击提交 会发请求到这里
			//1. 从请求中把数据拿出来
			var todo Todo
			//绑定
			if c.BindJSON(&todo) != nil {
				fmt.Println("c.BindJSON failed err =", err)
				return
			}
			//2. 存入数据库
			//3. 返回响应
			if db.Create(&todo).Error != nil {
				fmt.Println("db.Create failed err =", err)
				return
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})

		//查看所有待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			//查询todo这个表里所有的数据
			var todoList []Todo
			if db.Find(&todoList).Error != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})

		//查看某一条待办事项
		//v1Group.GET("/todo/:id", func(c *gin.Context) {
		//
		//})

		//删除某一条待办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			//拿到 id
			id, ok := c.Params.Get("id")
			//如果 id 不存在
			if !ok {
				c.JSON(http.StatusOK, gin.H{"err": "id not exist"})
				return
			}
			//删除
			if db.Where("id=?", id).Delete(Todo{}).Error != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})

		//修改某一条待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			//拿到 id
			id, ok := c.Params.Get("id")
			//如果 id 不存在
			if !ok {
				c.JSON(http.StatusOK, gin.H{"err": "id not exist"})
				return
			}
			//创建一个变量接收查询到的值
			var todo Todo
			if db.Where("id=?", id).First(&todo).Error != nil {
				c.JSON(http.StatusOK, gin.H{"err": err.Error()})
				return
			}
			//绑定
			if c.BindJSON(&todo) != nil {
				fmt.Println("c.BindJSON failed err =", err)
				return
			}
			//更新
			if db.Save(&todo).Error != nil {
				c.JSON(http.StatusOK, gin.H{"err": err.Error()})
				return
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
	}

	//访问 localhost:10824/bubble 之后会变成 localhost:10824/bubble#/
	//这是前端的事情，与路由配置有关
	err = r.Run(":10824")
	if err != nil {
		fmt.Println("run failed err =", err)
		return
	}
}
