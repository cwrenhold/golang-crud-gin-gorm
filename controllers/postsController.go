package controllers

import (
	"golang-crud-gin-gorm/initializers"
	"golang-crud-gin-gorm/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Create a new post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return the new post
	c.JSON(200, gin.H{
		"post": post,
	})
}
