package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func posts(c *gin.Context) {
	//这里叫 pp.tmpl 是因为 在 template/posts/index.tmpl 中定义了 {{define "pp.tmpl"}}
	c.HTML(200, "pp.tmpl", gin.H{
		"title": "www.baidu.com",
	})
}

func users(c *gin.Context) {
	c.HTML(200, "uu.tmpl", gin.H{
		//Gin 会转义，同样需要自定义函数
		"title": "<a href='https://www.baidu.com'>百度</a>",
	})
}

//静态文件：html 页面上用到的样式文件，如 css，js 等
//存储在 statics 文件夹里
func main() {
	r := gin.Default()

	//加载静态文件
	//前面的 /static 是默认叫这个名，就这么写是约定俗成的
	//后面的是相对于 main.go 同级的相对路径名，不用 ./
	r.Static("/static", "staticTest")

	r.SetFuncMap(template.FuncMap{
		"mySafe1": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//解析模板
	//两个模板要解析，可以这样写，但 20 个呢
	//r.LoadHTMLFiles("./templates/posts/index.tmpl", "./templates/users/index.tmpl")
	//20 个就需要这样写
	//Gin 里就不用 ./ 了
	//这里是通配符匹配
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", posts)
	r.GET("/users/index", users)

	//网上下载的模板
	//要注意对应 30 行改名
	r.GET("/internet", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})
	err := r.Run(":7777")
	if err != nil {
		return
	}
}
