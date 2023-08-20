package main

import "fmt"

func main() {
	//创建一个 int 型的 channel，容量为3
	myChan := make(chan int, 3)
	//写入 int 数据 3
	myChan <- 3
	//写入 int 数据 11
	myChan <- 11
	//写入 int 数据 2
	myChan <- 2

	//继续写入 int 数据 111，会报错
	//myChan <- 111

	//取数据
	n := <-myChan
	fmt.Println("取第1次", n) //结果为 3，因为 channel 本质上是一个队列
	n = <-myChan
	fmt.Println("取第2次", n) //结果为 11
	n = <-myChan
	fmt.Println("取第3次", n) //结果为 2

	//n = <-myChan
	//fmt.Println(n) //在没有使用协程的情况下，继续取会报错

	//此时管道为空，写入两个数据
	myChan <- 1000
	myChan <- 1001
	//关闭管道
	close(myChan)
	//此时不能继续写入，但仍可以读取
	n = <-myChan
	fmt.Println("关闭管道之后，取到", n) // 1000

	myChan2 := make(chan int, 100)
	for i := 1; i <= 100; i++ {
		myChan2 <- i * i
	}
	//遍历管道前，必须先关闭
	close(myChan2)
	//且只能用 for-range 遍历
	for v := range myChan2 {
		fmt.Printf("%d ", v)
	}
}
