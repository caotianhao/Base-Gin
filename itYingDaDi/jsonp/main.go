package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/jsonp", func(c *gin.Context) {
		//在 jsonp 格式下的域名后面加 ?callback=xxx
		//输出就会从 {"status":"ok"} 变成 xxx({"status":"ok"});
		//主要解决跨域问题
		c.JSONP(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	err := r.Run()
	if err != nil {
		return
	}
}
