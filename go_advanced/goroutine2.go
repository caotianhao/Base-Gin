package main

import (
	"fmt"
	"sync"
	"time"
)

//要求：计算 1,1+2,1+2+3,...,1+...+200 并把结果存到一个 map 中

var (
	myMap = map[int]int{}
	lock  sync.Mutex
)

func cal(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	//写的时候加锁
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	//fatal error: concurrent map writes
	//同时对 map 写，很严重的错误
	//需要使用锁或者 channel

	//for i := 1; i <= 200; i++ {
	//	go cal(i)
	//}
	//for i := 1; i <= 200; i++ {
	//	fmt.Printf("myMap[%d] = [%d]\n", i, myMap[i])
	//}

	//加锁后
	for i := 1; i <= 200; i++ {
		go cal(i)
	}
	//加锁之后要等待时间，否则读进程完了，map 还没写完，就会出现后面的值为 0
	time.Sleep(time.Second)
	//读的时候也加锁，为什么？
	//按理说 1 秒时间上面的协程都应该执行完，后面就不应该出现资源竞争的问题了
	//但是在实际运行中，还是可能出现（运行时增加 -race 参数，确实会发现有资源竞争问题)
	//因为我们程序从设计上可以知道 1 秒就执行完所有协程，但是主线程并不知道
	//因此底层可能仍然出现资源争夺，因此加入互斥锁即可解决问题
	lock.Lock()
	for i := 1; i <= 200; i++ {
		fmt.Printf("myMap[%d] = [%d]\n", i, myMap[i])
	}
	lock.Unlock()
	//但是这个程序可以等待 1 秒，其他程序到底是多少秒
	//这就需要使用管道通信来解决了
}
