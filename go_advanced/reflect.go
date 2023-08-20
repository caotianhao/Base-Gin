package main

import (
	"encoding/json"
	"fmt"
)

type cat struct {
	Name   string  `json:"Her name"`
	Age    int     `json:"Her age"`
	Weight float64 `json:"Her weight"`
}

func main() {
	myCat := cat{"jyz", 24, 90}
	data, _ := json.Marshal(myCat)
	fmt.Println(string(data)) //这里就用到了反射
}
