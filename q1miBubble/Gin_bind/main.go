package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//ShouldBind 会按照下面的顺序解析请求中的数据完成绑定：
//如果是 get 请求，只使用 form 绑定引擎（query）
//如果是 post 请求，首先检查 content-type 是否为 json 或 xml，然后再使用 form
type person struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}

func main() {
	r := gin.Default()
	r.GET("/lesson14", func(c *gin.Context) {
		//以 query 为例，字段少时我们可以一个一个用变量存储
		//但是字段太多这样就太麻烦了，所以应该用 gin 的参数绑定
		//先初始化一个结构体变量
		var p person
		//使用 ShouldBind 来进行绑定，因为要改变 p 的值，所以是引用类型
		//ShouldBind 不止可以对 query 数据进行处理，json 和 form 都可以
		err := c.ShouldBind(&p)
		if err != nil {
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"name": p.Name,
				"age":  p.Age,
			})
		}
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
