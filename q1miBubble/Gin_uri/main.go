package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取请求的 path(uri) 参数
func main() {
	r := gin.Default()
	//这里假设我再写一个，比如查某年某月的博客
	//如果写成 r.GET("/:year/:month", func(c *gin.Context){} )
	//就会出现路由匹配错误，因此最好在前面加上区分
	//写成 r.GET("/blog/:year/:month", func(c *gin.Context){} )
	//这样访问时就是 :8080/lesson13 和 :8080/blog 两个页面了
	r.GET("/lesson13/:name/:age", func(c *gin.Context) {
		//例如，我访问的是 localhost:8080/lesson13/bob/15
		//则 /:name/:age 就对应的接收了
		//并且可以使用 c.Param 赋给相应的值来使用
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
