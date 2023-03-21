package middleware

import (
	"context"
	"log"
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
			ctx := context.Background()
			cachedData, err := config.Client.Get(ctx, cacheKey).Result()
			if err != nil {
				c.Next()

				// If the request was successful, cache the data
				if c.Writer.Status() == http.StatusOK {
					data := c.MustGet("data").([]byte)

					if url == "/posts/:id" && config.CACHE_POST {
						cacheErr := config.Client.Set(ctx, cacheKey, data, 0).Err()
						if cacheErr != nil {
							log.Println("🚀 Error caching data")
							c.Abort()
							return
						}
					} else if url == "/todos/:id" && config.CACHE_TODO {
						cacheErr := config.Client.Set(ctx, cacheKey, data, 0).Err()
						if cacheErr != nil {
							log.Println("🚀 Error caching data")
							c.Abort()
							return
						}
					}
				}
				return
			}

			c.Data(http.StatusOK, "application/json", []byte(cachedData))
			c.Abort()
			return
		}
	}
}
