package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age  int
}

func reflectDemo(i interface{}) {
	rTyp := reflect.TypeOf(i)
	rVal := reflect.ValueOf(i)
	fmt.Println("rTyp =", rTyp, ", rVal =", rVal)

	//rVal 并不是 int 型
	fmt.Printf("rVal type is %T\n", rVal)

	//ans := 1 + rVal //会提示 mismatched types int and reflect.Value
	ans := 1 + rVal.Int() //int64
	fmt.Println("1 + rVal.Int() =", ans)

	//int, int
	fmt.Printf("rVal kind is %v, rVal type is %v\n", rVal.Kind(), rVal.Type())
}

func reflectDemoStruct(i interface{}) {
	rTyp := reflect.TypeOf(i)
	rVal := reflect.ValueOf(i)
	fmt.Println("rTyp =", rTyp, ", rVal =", rVal)

	// 转成接口 --> 类型断言 --> 取其中一项
	fmt.Println("rVal.Interface().(person).name =", rVal.Interface().(person).name)

	//struct, main.person(包名.person)
	fmt.Printf("rVal kind is %v, rVal type is %v\n", rVal.Kind(), rVal.Type())

	//如果以后遇见很多结构体都有 name 字段，上面直接转 person 易错
	//可以使用 switch + x.(type) 判断
	//rVal.Interface().(type) 和 rVal.Type() 不一样
	switch rVal.Interface().(type) {
	case int:
		fmt.Println("int")
	case person:
		fmt.Println("rVal.Interface().(type) = person")
	}
}

func main() {
	n := 7
	reflectDemo(n)
	fmt.Println("----------------------------------------------")
	p := person{"alice", 18}
	reflectDemoStruct(p)
}
