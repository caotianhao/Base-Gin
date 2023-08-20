package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Student struct {
	Name  string `json:"name"`
	Age   int
	Score float64
}

func main() {
	r := gin.Default()
	s := Student{"Alice", 90, 22.33}
	r.GET("/", func(c *gin.Context) {
		c.XML(http.StatusOK, s)
	})
	err := r.Run()
	if err != nil {
		return
	}
}
