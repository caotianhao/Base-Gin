package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//1. 定义模板，见 hello.tmpl
		//2. 解析模板
		files, err := template.ParseFiles("./hello.tmpl")
		if err != nil {
			fmt.Println("template.ParseFiles err =", err)
			return
		}
		//3. 渲染模板
		err = files.Execute(w, "cth")
		if err != nil {
			fmt.Println("files.Execute err =", err)
			return
		}
	})
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("ListenAndServe err =", err)
		return
	}
}
