package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to CacheRocket! The server is running!",
	})
}
