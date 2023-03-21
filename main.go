package main

import (
	"github.com/anshg1214/CacheRocket/config"
	"github.com/anshg1214/CacheRocket/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Routes
	router.GET("/posts/:id", controller.GetPost)
	router.GET("/todos/:id", controller.GetTodo)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey! The server is running!",
		})
	})

	err := router.Run(":" + config.PORT)

	if err != nil {
		panic(err)
	}
}
