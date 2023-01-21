package main

import (
	"golang-crud-gin-gorm/controllers"
	"golang-crud-gin-gorm/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/posts", controllers.PostsCreate)

	r.Run() // listen and serve on 0.0.0.0:{PORT}
}
