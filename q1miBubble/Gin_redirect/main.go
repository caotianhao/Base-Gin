package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//1. http 重定向
	//我访问 localhost:8080/lesson16，我直接转移到新的网站
	//如果遇到直接把新网址加到 localhost:8080 后面的，清理浏览器缓存就解决了
	r.GET("/lesson16", func(c *gin.Context) {
		//这样处理地址栏的地址变了，完全跳转过去了相当
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	//2. 路由重定向
	//这样处理地址栏就没变，地址栏仍然是 localhost:8080/a，而内容转向显示 b 的
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"A OR B": "This is B.",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
