package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/anshg1214/CacheRocket/config"
	"github.com/anshg1214/CacheRocket/utils"
	"github.com/gin-gonic/gin"
)

var cacheLock sync.Map

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

			// Check if data is cached
			cachedData, err := utils.GetValueFromCache(cacheKey, ctx)
			if err == nil {
				c.Data(http.StatusOK, "application/json", []byte(cachedData))
				c.Abort()
				return
			}

			// Data is not cached, acquire a lock for this key
			lock, _ := cacheLock.LoadOrStore(cacheKey, &sync.Mutex{})
			lock.(*sync.Mutex).Lock()
			defer lock.(*sync.Mutex).Unlock()

			// Check if another request has fetched the data while we were waiting for the lock
			cachedData, err = utils.GetValueFromCache(cacheKey, ctx)
			if err == nil {
				c.Data(http.StatusOK, "application/json", []byte(cachedData))
				c.Abort()
				return
			}

			c.Next()

			// If the request was successful, cache the data
			if c.Writer.Status() == http.StatusOK {
				data := c.MustGet("data").([]byte)

				if url == "/posts/:id" && config.CACHE_POST {
					cacheErr := utils.SetDataInCache(cacheKey, data, ctx)
					if cacheErr != nil {
						log.Println("🚀 Error caching data")
						c.Abort()
						return
					}
				} else if url == "/todos/:id" && config.CACHE_TODO {
					cacheErr := utils.SetDataInCache(cacheKey, data, ctx)
					if cacheErr != nil {
						log.Println("🚀 Error caching data")
						c.Abort()
						return
					}
				}
			}
			return
		}
	}
}
