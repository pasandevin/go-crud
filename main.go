package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pasandevin/go-crud/controllers"
	"github.com/pasandevin/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)

	r.Run()
}
