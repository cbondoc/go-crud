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

	//Create post
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

func PostsIndex(c *gin.Context) {
	// https://gorm.io/docs/query.html
	// Retrieving all objects
	// // Get all records
	// result := db.Find(&users)
	// // SELECT * FROM users;
	// result.RowsAffected // returns found records count, equals `len(users)`
	// result.Error        // returns error

	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id of URL
	id := c.Param("id")

	// db.First(&user, 10)
	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id off URL
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	// Based on the documentation of gorm https://gorm.io/docs/update.html
	// 	 Update attributes with `struct`, will only update non-zero fields
	// db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
	//  UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;
	// Update attributes with `struct`, will only update non-zero fields
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get id off the URL
	id := c.Param("id")

	// Delete it
	initializers.DB.Delete(&models.Post{}, id)

	// Respond with it
	c.Status(200)
}
