package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)


func main() {
	// e := gin.New()
	// _ = e
	r := gin.Default()

	r.StaticFile("/", "./public")

	r.GET("/ping", func(c *gin.Context) {
		foo := c.Query("foo")
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("pong %s", foo),
		})
	})

	r.GET("/foo/:bar", func(c *gin.Context) {
		val := c.Param("bar")
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("foo %s", val),
		})
	})

	// r.POST("/login")

	r.Run() // listen and serve on 0.0.0.0:8080
}
