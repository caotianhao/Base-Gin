package main

import (
	"html/template"
	"net/http"
)

func nestDemo(w http.ResponseWriter, r *http.Request) {
	//这里一定要把被包含的写在后面
	t, err := template.ParseFiles("./nest.tmpl", "./ul.tmpl")
	if err != nil {
		return
	}
	name := "cth nest test"
	err = t.Execute(w, name)
	if err != nil {
		return
	}
	//输出以下内容：
	//测试嵌套
	//ul 1
	//ul 2
	//ul 3
	//define 1
	//define 2
	//define 3
	//hello cth nest test
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//内置函数不够了，我想自定义函数
		//这里自定义函数要么只有一个返回值，要么有两个且第二个返回值必须为 error
		kua := func(name string) (string, error) {
			return name + " is handsome", nil
		}

		//与前两个不同的是，这里不是直接使用 template.ParseFiles
		//而是先创建了一个名字为 mb3 的模板对象
		//New 的名字一定要和使用的模板名字相同，而且格式也要带上
		//假设 ParseFiles 传入不止一个模板，那么 New 的那个一定是其中之一才行
		t := template.New("mb3.tmpl")

		//一定要在解析模板之前，告诉我自己定义了一个自定义函数 kua
		//"kuaKua": kua 前面的 kuaKua 是在模板中使用的名字，后面的是为了建立联系
		t.Funcs(template.FuncMap{"kuaKua": kua})

		//然后再解析模板
		t, err := t.ParseFiles("./mb3.tmpl")
		if err != nil {
			return
		}
		name := "cth"
		err = t.Execute(w, name)
		if err != nil {
			return
		}
	})
	http.HandleFunc("/nestDemo", nestDemo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
