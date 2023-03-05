package main

import (
	"gin-api/project/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})
	db.Init()
	err := r.Run()
	if err != nil {
		return
	}
	db.Close()
}
