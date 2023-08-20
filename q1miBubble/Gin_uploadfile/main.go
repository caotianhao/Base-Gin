package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/lesson15", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/lesson15", func(c *gin.Context) {
		//从请求中读取文件
		f, err := c.FormFile("file1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		} else {
			//将读取到的文件保存到本地（即服务端）
			//使用路径拼接，传入到当前文件夹下，测试成功
			dst := path.Join("./", f.Filename)
			err := c.SaveUploadedFile(f, dst)
			if err != nil {
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": "ok",
				})
			}
		}
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
