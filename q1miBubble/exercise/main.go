package main

import (
	"html/template"
	"net/http"
)

type student struct {
	Name string
	Age  int
}

func exercise(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./temp.tmpl")
	if err != nil {
		return
	}
	s1 := "1111"
	s2 := 188
	s3 := 'r'
	s4 := student{"cao", 25}
	err = t.Execute(w, map[string]interface{}{
		"s11": s1,
		"s22": s2,
		"s33": string(s3),
		"s44": s4,
	})
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/exercise", exercise)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
