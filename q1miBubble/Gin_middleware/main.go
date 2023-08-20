package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//Gin 中间件
//Gin 框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数
//这个钩子函数就叫中间件，适合处理一些公共的业务逻辑
//比如登录认证，权限校验，数据分页，记录日志，耗时统计等
//Gin 的中间件必须是 gin.HandlerFunc 型的
//注意到平时使用的 r.GET("/lesson18", func(c *gin.Context) {}
//其中的第二个参数就是 gin.HandlerFunc 型的，这其实就是一个中间件
//不使用匿名函数定义的话，单独拿出来就像是自己定义的中间件，例如：
func lesson18zjj(c *gin.Context) {
	fmt.Println("---lesson18zjj begin---")
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
	fmt.Println("---lesson18zjj end---")
}

//定义一个计算耗时的中间件
func calTime(c *gin.Context) {
	fmt.Println("---calTime begin---")
	start := time.Now()

	c.Next() //调用后续的处理函数（中间件）
	//c.Abort() //不调用后续的处理函数（中间件）

	cost := time.Since(start)
	fmt.Println("cost time:", cost)
	fmt.Println("---calTime end---")
}

func m1(c *gin.Context) {
	fmt.Println("---m1 begin---")
	c.Next()
	//c.Abort()
	//即使是 c.Abort() 也会输出 m1 end，要注意
	//若是不想输出 m1 end，则应 c.Abort() 之后 return
	//return
	fmt.Println("---m1 end---")
}

func m2(c *gin.Context) {
	c.Set("name", "alice")
}

/*
//中间件的应用------------------------------------------------------------------------------------
//比如一个登录系统
//是否为登录用户
func isUser() bool {
	return true
}

func authMiddleWare(c *gin.Context) {
	//验证是否为登录用户
	if isUser() {
		c.Next()
	} else {
		c.Abort()
	}
}

//但是经常使用闭包的形式：
//这个 (check bool) 就是标记是否要进行登录检查
//之后使用 r.Use(authMiddleWareNew(true)) 表示进行登录检查
func authMiddleWareNew(check bool) gin.HandlerFunc {
	//写成闭包是为了连接数据库或进行其他准备工作等等
	return func(c *gin.Context) {
		if isUser() {
			c.Next()
		} else {
			c.Abort()
		}
	}
}
//中间件的应用------------------------------------------------------------------------------------
*/

func main() {
	//中间件注意事项：
	//gin.Default() 默认使用了 Logger 和 Recover 中间件，其中
	//Logger 中间件将日志写入 gin.DefaultWriter，即使配置了 GIN_MODE=release
	//Recover 中间件会 recover 任何 panic，如果有 panic 的话，会报 500 响应码
	//如果不想使用上面两个默认的中间件，可以使用 gin.New() 新建一个没有任何默认中间件的路由
	r := gin.Default()

	//中间件的另一个常用用法：跨中间件存取值
	//在中间件中定义一个 key-value，然后在 get 请求里获取
	r.GET("/middleWare", m2, func(c *gin.Context) {
		//可以使用 c.Get 获取，若 key 存在，则会返回对应的 value 和 true
		name, ok := c.Get("name")
		if !ok {
			name = "No name"
		}
		//也可以使用
		//name := c.MustGet("name")
		//但这样的话，key 不存在的话就会 panic
		c.JSON(http.StatusOK, gin.H{"name": name})
	})

	r.GET("/lesson18", calTime, lesson18zjj)
	//输出以下 6 行
	/*
		---calTime begin---
		---lesson18zjj begin---
		---lesson18zjj end---
		cost time: 519.9µs
		---calTime end---
		[GIN] 2023/03/06 - 14:44:49 | 200 | 519.9µs | 127.0.0.1 | GET "/lesson18"
	*/

	//假设不仅仅访问 /lesson18 需要 calTime 中间件
	//例如访问 /index，/home，/shop 等都需要
	//这时可以定义一/多个全局中间件
	r.Use(calTime, m1)
	//此时接下来的所有访问请求都要经过该一/多个中间件
	r.GET("/index", func(c *gin.Context) {
		fmt.Println("index")
		c.JSON(http.StatusOK, gin.H{"msg": "index"})
	})
	r.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "home"})
	})
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "shop"})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}

	//此外，当在中间件或 Handler 中启动新的 goroutine 时
	//不能使用原始的上下文 (c *gin.Context)，只能使用其只读副本 (c.Copy())
}
