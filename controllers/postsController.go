package controllers

import (
	"github.com/cbondoc/go-crud/initializers"
	"github.com/cbondoc/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Get data off request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return

	}
	//Return it

	c.JSON(200, gin.H{
		"post": post,
	})

}
