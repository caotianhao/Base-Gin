package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//访问 /lesson17 的 get 请求会走这一条逻辑
	r.GET("/lesson17", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "get",
		})
	})

	//访问 /lesson17 的 post 请求会走这一条逻辑
	r.POST("/lesson17", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "post",
		})
	})

	//访问 /lesson17 的 put 请求会走这一条逻辑
	r.PUT("/lesson17", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "put",
		})
	})

	//访问 /lesson17 的 delete 请求会走这一条逻辑
	r.DELETE("/lesson17", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "delete",
		})
	})

	//上面那样写比较麻烦
	//可以使用 r.Any 配合 switch-case
	//注意这里就不要写 /lesson17 了，会提示已占用
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{"method": "get"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "post"})
		case http.MethodPut:
			c.JSON(http.StatusOK, gin.H{"method": "put"})
		case http.MethodDelete:
			c.JSON(http.StatusOK, gin.H{"method": "delete"})
		}
	})

	//上面我定义了个 /lesson17 和 /user
	//但是，用户什么都有可能访问，代码层面显然不可能全都定义到
	//而又不希望访问不存在页面的时候转到 404
	//这时，应该用 r.NoRoute
	//自然地，这里面就没有地址，只有匿名函数
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "404!",
		})
	})

	//路由组
	//我一个 video 页面下面细分了很多
	//这样一个一个写很麻烦，冗长
	//所以就有了路由组的写法
	//r.GET("/video/buy", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"This is": "video/buy"})
	//})
	//r.GET("/video/watch", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"This is": "video/watch"})
	//})
	//r.GET("/video/download", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"This is": "video/download"})
	//})
	//这里大括号加是为了看起来有条理，清晰，不加也可以
	//路由组也是支持嵌套的，多用于区分业务逻辑以及 API 版本等
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/buy", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"This is": "video/buy"})
		})
		videoGroup.GET("/watch", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"This is": "video/watch"})
		})
		videoGroup.GET("/download", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"This is": "video/download"})
		})
	}

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
