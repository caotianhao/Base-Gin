package main

import (
	"html/template"
	"net/http"
)

//修改标识符
//因为在 vue 等前端里面标识符也是 {{}}, 为了避免冲突要修改模板引擎的标识符
//例子：{{}} 修改为 {[]}
func lesson8(w http.ResponseWriter, r *http.Request) {
	//在解析模板 (ParseFiles) 之前
	//t, err := template.ParseFiles("./lesson8.tmpl")
	t, err := template.New("lesson8.tmpl").Delims("{[", "]}").ParseFiles("./lesson8.tmpl")
	//这样就可以使用 {[]} 了，同时不能使用 {{}} 了
	if err != nil {
		return
	}
	name := "标识"
	err = t.Execute(w, name)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/lesson8", lesson8)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
