package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取 form 表单提交的数据
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./lesson12.html", "./index.html")
	r.GET("/lesson12", func(c *gin.Context) {
		c.HTML(http.StatusOK, "lesson12.html", nil)
	})
	r.POST("/lesson12", func(c *gin.Context) {
		//第 1 种方法
		//username := c.PostForm("username")
		//password := c.PostForm("password")

		//第 2 种方法
		//username := c.DefaultPostForm("username", "nobody")
		//这里假设密码不填，会认为你的 password 为空，从而返回的不是 ******
		//因为这里的 default 指的是取不到这个字段，应该这样写才是 ******，因为没有 xxx 字段
		//password := c.DefaultPostForm("xxx", "******")
		//password := c.DefaultPostForm("password", "******")

		//第 3 种方法，同样类似于 query
		username, okU := c.GetPostForm("username")
		if !okU {
			username = "nobody"
		}
		password, okP := c.GetPostForm("password")
		if !okP {
			password = "******"
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"username": username,
			"password": password,
		})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
