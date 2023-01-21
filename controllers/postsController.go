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

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get the id from the URL
	id := c.Param("id")

	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id from the URL
	id := c.Param("id")

	// Get the data from request body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Find the post to update
	var post models.Post
	initializers.DB.First(&post, id)

	// Update the post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with the updated post
	c.JSON(200, gin.H{
		"post": post,
	})
}
