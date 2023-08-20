package main

import (
	"html/template"
	"net/http"
)

func xss(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"mySafe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		return
	}
	//html/template 和 text/template 的区别：
	//html/template 会对有风险的内容全部进行转义
	//这里输入 xss 内容
	msg1 := "<script>alert(123);</script>"
	//但有些时候不想全都转义，例如 msg2
	//同样，这就需要在解析模板之前 new 一个
	//自定义一个名字为 mySafe 的函数
	msg2 := "<a href='https://www.baidu.com'>百度哦</a>"
	//这里传一个 map 进去，一定要记得修改模板文件里的
	//把 <p>{{.}}</p> 改为 <p>{{.msg1}}</p> <p>{{.msg2}}</p>
	//这样输出的两个仍然都是转义之后的
	//如果想转义 msg1，而保留 msg2
	//则应把 <p>{{.msg2}}</p> 改为 <p>{{.msg2|mySafe}}</p>
	//这样的话就会只保留“百度哦”三个字，并且是个链接
	err = t.Execute(w, map[string]string{
		"msg1": msg1,
		"msg2": msg2,
	})
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
