package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("-----hello test", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //开启一个协程，可以满足同时打印，主线程结束，协程也结束
	//test() //这样是先执行 test 之后再执行 main 的
	for i := 1; i <= 10; i++ {
		fmt.Println("hello main", i)
		time.Sleep(time.Second / 3)
	}
}
