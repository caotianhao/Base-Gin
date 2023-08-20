package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"exe2age"`
	Score float64
}

func main() {
	route := gin.Default()

	p1 := Person{"alice", 12, 123.45}
	p2 := Person{"bob", 22, 456.78}
	route.GET("/exe2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"p1Age": p1.Age,
			"p1":    p1,
			"p2":    p2,
		})
	})
	err := route.Run(":9999")
	if err != nil {
		return
	}
}
