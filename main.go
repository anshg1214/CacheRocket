package main

import (
	"log"

	"github.com/anshg1214/CacheRocket/config"
	"github.com/anshg1214/CacheRocket/controller"
	"github.com/anshg1214/CacheRocket/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Routes
	router.GET("/", controller.GetHome)
	router.GET("/posts/:id", middleware.CacheMiddleware(), controller.GetPost)
	router.GET("/todos/:id", middleware.CacheMiddleware(), controller.GetTodo)

	redisErr := config.PingRedis()
	if redisErr != nil {
		log.Fatal("❌ Redis is not running")
	}

	err := router.Run(":" + config.PORT)
	if err != nil {
		panic(err)
	}
}
