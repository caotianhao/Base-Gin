package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取浏览器那边发来请求携带的 query string 参数
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./lesson11.tmpl")
	r.GET("/lesson11", func(c *gin.Context) {
		//1. 访问时地址栏需要输入 localhost:8080/lesson11?query=xxx
		//name := c.Query("query")

		//2. 还可以使用 DefaultQuery，查到的时候就用查到的，查不到就用 Default 的
		//name := c.DefaultQuery("query", "None")

		//3. 还可以使用 GetQuery，前两个组合起来
		//name, ok := c.GetQuery("query")
		//if !ok {
		//	name = "None"
		//}

		//访问多个，中间用 & 连接，例如 ?name=alice&age=18
		name := c.Query("name")
		age := c.Query("age")
		c.HTML(http.StatusOK, "lesson11.tmpl", gin.H{
			"name": name,
			"age":  age,
		})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
