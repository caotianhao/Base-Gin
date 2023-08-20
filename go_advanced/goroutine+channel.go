package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("write success", i)
		time.Sleep(time.Second / 25)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		i, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("read success", i)
		time.Sleep(time.Second / 25)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	//exitChan <- true

	//开启协程，交替读写
	go writeData(intChan)
	go readData(intChan, exitChan)

	//这样就卡住了，不用人为的估计需要 sleep 多少秒
	for {
		_, ok := <-exitChan
		//这里为什么是 !ok
		//因为没有一个直接的方式来判断是否通道已经关闭，但是这里有接收操作的一个变种
		//它产生两个结果：接收到的通道元素，以及一个布尔值（通常称为 ok）
		//true 的时候代表接收成功
		//false 表示当前的接收操作在一个关闭的并且读完的通道上
		//所以这里是判断 exitChan 是否已空且关闭
		if !ok {
			break
		}
	}
}
