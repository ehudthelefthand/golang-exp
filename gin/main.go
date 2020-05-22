package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
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

	r.POST("/login", func (c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(200, gin.H{
			"username": username,
			"password": password,
		})
	})

	r.POST("/login2", func (c *gin.Context) {
		type User struct {
			Username string 
			Password string
		}
		user := User{}
		err := c.BindJSON(&user)
		if err != nil {
			c.Status(400)
			return
		}
		log.Printf("%v\n", user)
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.POST("/login3", func (c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		form, err := c.MultipartForm()
		if err != nil {
			c.Status(400)
			return
		}
		log.Printf("%v\n", username)
		log.Printf("%v\n", password)
		log.Printf("%v\n", form.File["photo"])
		file := form.File["photo"][0]
		err = c.SaveUploadedFile(file, 
			"./upload/" + file.Filename)
		if err != nil {
			c.Status(400)
			return
		}
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
