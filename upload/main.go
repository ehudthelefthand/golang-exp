package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// r.Static("/image", "./dest")
	r.Static("/", "./public")

	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println(err)
			c.Status(400)
			return
		}

		files := form.File["photo"]
		for _, file := range files {
			c.SaveUploadedFile(file, file.Filename)
			// src, err := file.Open()
			// fmt.Println(err)
			// defer src.Close()
			// dst := filepath.Join("dest", file.Filename)
			// out, err := os.Create(dst)
			// fmt.Println(err)
			// defer out.Close()
			// _, err = io.Copy(out, src)
			// fmt.Println(err)
		}
	})

	r.Run(":9090")
}
