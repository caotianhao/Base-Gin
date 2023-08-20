package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/exerciseGin", func(c *gin.Context) {
		c.String(200, "what")
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
