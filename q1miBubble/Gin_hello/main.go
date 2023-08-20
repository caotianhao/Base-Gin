package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//参数也是定死的
//func sayHelloGin(c *gin.Context) {
//	//200 是 http 状态码，也可以使用 http.StatusOK 代替
//	c.JSON(200, gin.H{
//		"message": "helloGin",
//	})
//}

func main() {
	//返回默认的路由引擎
	r := gin.Default()

	//指定用户使用 GET 请求访问 /helloGin
	//函数形式
	//r.GET("/helloGin", sayHelloGin)

	//匿名函数形式
	r.GET("/helloGin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})

	//启动服务
	err := r.Run(":9090")
	if err != nil {
		return
	}
}
