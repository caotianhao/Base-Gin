package main

import "fmt"

func main() {
	const (
		n1 = iota //表示赋值为0
		n2        //在 n1 的基础上 +1
		n3        //+1..以此类推
	)
	fmt.Println(n1, n2, n3) //0,1,2

	const (
		n4     = iota
		n5, n6 = 7, 9
	)
	fmt.Println(n4, n5, n6) //0,7,9

	const (
		n7      = iota
		n8      = iota
		n9, n10 = iota, iota
		n11     = iota
	)
	fmt.Println(n7, n8, n9, n10, n11) //0,1,2,2,3
}
