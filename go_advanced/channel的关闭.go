package main

import "fmt"

func main() {
	chan1 := make(chan int, 3)
	chan1 <- 1
	n1, ok1 := <-chan1
	fmt.Println("n1 =", n1, "管道未空且未关闭", ok1) //true

	//管道已空且未关闭，报错
	//_, ok2 := <-chan1
	//fmt.Println("管道已空且未关闭", ok2)

	chan1 <- 2
	close(chan1)
	n3, ok3 := <-chan1
	fmt.Println("n3 =", n3, "管道未空且已关闭", ok3) //true

	n4, ok4 := <-chan1
	fmt.Println("n4 =", n4, "管道已空且已关闭", ok4) //false
}
