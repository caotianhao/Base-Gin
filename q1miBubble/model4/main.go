package main

import (
	"net/http"
	"text/template"
)

//很多网页长得都差不多，这时候就需要模板继承
func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		return
	}
	msg := "cth"
	err = t.Execute(w, msg)
	if err != nil {
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		return
	}
	msg := "cth"
	err = t.Execute(w, msg)
	if err != nil {
		return
	}
}

//这里把 ./base.tmpl 写成 ./base,tmpl 了......改了半天
func indexNew(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./base.tmpl", "./indexNew.tmpl")
	if err != nil {
		return
	}
	msg := "cth new 试试 ExecuteTemplate 不用行不行"
	//这里渲染模板使用的是 ExecuteTemplate, 因为 base 继承了不止一个模板
	//试过了，不用也行
	err = t.ExecuteTemplate(w, "indexNew.tmpl", msg)
	if err != nil {
		return
	}
}

func homeNew(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./base.tmpl", "./homeNew.tmpl")
	if err != nil {
		return
	}
	msg := "cth new"
	err = t.ExecuteTemplate(w, "homeNew.tmpl", msg)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/indexNew", indexNew)
	http.HandleFunc("/homeNew", homeNew)
	//访问 localhost:8080/index 和 localhost:8080/home 发现长得太像了
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
