package main

import (
	"html/template"
	"net/http"
)

type student struct {
	Name  string
	Age   int
	Score float64
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		files, err := template.ParseFiles("./hello.tmpl")
		if err != nil {
			return
		}
		//上次实验中，传了字符串进去，如下：
		//err = files.Execute(w, "ddd")
		//这次想要传结构体
		//stu1 := student{
		//	Name:  "cth",
		//	Age:   25,
		//	Score: 100.0,
		//}
		//err = files.Execute(w, stu1)
		//{{.}} --> Hi {cth 25 100}
		//----------------------------------------
		//<p>姓名 {{.Name}}</p>
		//<p>年龄 {{.Age}}</p>
		//<p>成绩 {{.Score}}</p>
		//姓名 cth
		//年龄 25
		//成绩 100

		//渲染模板传多个变量，同时需要修改模板
		s1 := student{
			Name:  "s1",
			Age:   100,
			Score: 56.7,
		}
		m1 := map[string]interface{}{
			"name":     "m1",
			"age":      90,
			"allScore": []int{1, 2, 3},
			"sign":     "哎嗨",
		}
		like := []string{"play", "do", "watch"}
		err = files.Execute(w, map[string]interface{}{
			"s1":   s1,
			"m1":   m1,
			"like": like,
		})
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
