package utils

import (
	"context"
	"net/http"

	"github.com/anshg1214/CacheRocket/config"
	"github.com/gin-gonic/gin"
)

func CacheData(cacheKey string, data []byte) error {
	ctx := context.Background()
	err := config.Client.Set(ctx, cacheKey, data, 0).Err()

	return err
}

func ThrowInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func ThrowInvalidJSONError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
}
