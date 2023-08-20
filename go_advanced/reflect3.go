package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string
	Age  int
}

//通过反射修改值
func reflectModify(i interface{}) {
	rVal := reflect.ValueOf(i)
	switch rVal.Interface().(type) {
	case *int:
		//不能使用 rVal.SetInt(20), 因为 SetInt 方法绑定的是非指针型
		//而传入的参数是 &n, 此时 rVal 是个指针类型的值
		//需要使用 Elem() 转换
		//就类似于使用 * 取到指针指向的值一样
		//fmt.Printf("%T\n", rVal.Elem())
		rVal.Elem().SetInt(20)
	case *string:
		rVal.Elem().SetString("no")
	case *student:
		//修改 Name 字段
		//Field(0) 表示结构体的第 0 个字段
		rVal.Elem().Field(0).SetString("Bob Field(0)")
		//FieldByName("Name") 表示 Name 字段
		rVal.Elem().FieldByName("Name").SetString("Bob FieldByName(\"Name\")")
		rVal.Elem().FieldByName("Age").SetInt(88)
	}
}

func main() {
	n := 10
	fmt.Println("before change:", n) //10
	reflectModify(&n)
	fmt.Println("after change:", n) //20
	fmt.Println("------------------------------------------")

	str := "yes"
	fmt.Println("before change:", str) //"yes"
	reflectModify(&str)
	fmt.Println("after change:", str) //"no"
	fmt.Println("------------------------------------------")

	stu1 := student{"bob", 30}
	fmt.Println("before change: he is", stu1.Name) //bob
	fmt.Println("before change: he is", stu1.Age)  //30
	reflectModify(&stu1)
	fmt.Println("after change: he is", stu1.Name) //Bob
	fmt.Println("after change: he is", stu1.Age)  //88
}
