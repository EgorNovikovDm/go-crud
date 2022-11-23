package controllers

import (
	"github.com/EgoriyNovikovDm/go-crud/initializers"
	"github.com/EgoriyNovikovDm/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//GET DATA OF REQUEST BODY

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//CREATE A POST

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//RETURN IT

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	//Respond with them

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get id from URL
	id := c.Param("id")

	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)
	//Respond with them

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get the id off the url

	id := c.Param("id")

	// Get the data off request body

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post were updating

	var post models.Post
	initializers.DB.First(&post, id)

	//Update id
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	//Get id off the url
	id := c.Param("id")
	//Delete post
	initializers.DB.Delete(&models.Post{}, id)

	//Respond
	c.Status(200)
}
