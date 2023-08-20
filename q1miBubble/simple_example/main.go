package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//规定必须接收两个参数 w http.ResponseWriter, r *http.Request
//包含了请求和响应
func sayHello(w http.ResponseWriter, r *http.Request) {
	//简单输出
	//_, _ = fmt.Fprintln(w, "hello golang Gin")

	//调整字体
	//_, _ = fmt.Fprintln(w, "<h1>hello golang Gin</h1>")//调整字体

	//文本太长，就放到文件里，比如 .txt
	//里面的东西就是 h5+css+js 啦
	file, err := ioutil.ReadFile("D:\\CTH\\GoProject\\src\\Gin\\q1miBubble\\简单入门\\hello.txt")
	if err != nil {
		fmt.Println("read file err =", err)
		return
	}
	//返回的 file 是 []byte 型，一定要转为 string
	_, _ = fmt.Fprintln(w, string(file))
}

func main() {
	//路径 + 函数
	//访问 /hello, 执行 sayHello 函数
	//若访问的不是 /hello, 而是不存在的网站，比如 /hello1, 会显示 404 page not found
	http.HandleFunc("/hello", sayHello)
	//仅这样还没有启动服务

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http sever failed with %v\n", err)
		return
	}
	//运行后访问 127.0.0.1:9090/hello 即可显示 hello golang Gin
	//这就是一个简单的网站
}
