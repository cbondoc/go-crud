package main

import (
	"github.com/cbondoc/go-crud/controllers"
	"github.com/cbondoc/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

// this will run before func main
func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.Run()
}
