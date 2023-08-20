package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type student struct {
	Name string `json:"name"`
	Age  int
}

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//Gin 渲染 JSON
		//方法 1：map 或 gin.H
		//data := map[string]interface{}{
		//	"name":    "cth",
		//	"age":     25,
		//	"message": "hello world",
		//}

		//方法 2：结构体
		data := student{"cth", 18}
		c.JSON(http.StatusOK, data)
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
