package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/anshg1214/CacheRocket/config"
	"github.com/gin-gonic/gin"
)

func CacheMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ID"})
			c.Abort()
			return
		}

		url := c.FullPath()
		cacheKey := strings.Split(url, "/")[1] + ":" + id

		if url == "/posts/:id" && !config.CACHE_POST {
			c.Next()
			return
		} else if url == "/todos/:id" && !config.CACHE_TODO {
			c.Next()
			return
		} else {
			fmt.Println("🚀 Fetching data from cache...")

			ctx := context.Background()
			cachedData, err := config.Client.Get(ctx, cacheKey).Result()
			if err != nil {
				c.Next()
				return
			}

			fmt.Println("🚀 Data fetched from cache")
			c.Data(http.StatusOK, "application/json", []byte(cachedData))
			c.Abort()
			return
		}
	}
}
