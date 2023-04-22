package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pasandevin/go-crud/initializers"
	"github.com/pasandevin/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	//get data off the body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	//create post
	post := models.Post{Title: body.Title, Body: body.Body}

	//save post to db
	result := initializers.DB.Create(&post)

	//check for errors
	if result.Error != nil {
		c.Status(400)
		return
	}

	//respond with post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//declare posts variable
	var posts []models.Post

	//get all posts
	initializers.DB.Find(&posts)

	//respond with posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//get id from url
	id := c.Param("id")

	//declare post variable
	var post models.Post

	//get post
	initializers.DB.First(&post, id)

	//respond with post
	c.JSON(200, gin.H{
		"post": post,
	})
}
