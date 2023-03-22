package utils

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/anshg1214/CacheRocket/config"
)

func CacheData(cacheKey string, data []byte) error {
	ctx := context.Background()
	err := config.Client.Set(ctx, cacheKey, data, 0).Err()

	return err
}

func ThrowInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	c.Abort()
}

func ThrowInvalidJSONError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
	c.Abort()
}

func ThrowNotFoundError(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{})
	c.Abort()
}

func GetValueFromCache(cacheKey string, ctx context.Context) ([]byte, error) {
	cachedData, err := config.Client.Get(ctx, cacheKey).Result()
	if err != nil {
		return nil, err
	}

	return []byte(cachedData), nil
}

func SetDataInCache(cacheKey string, data []byte, ctx context.Context) error {
	err := config.Client.Set(ctx, cacheKey, data, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func FlushCache() error {
	ctx := context.Background()
	err := config.Client.FlushAll(ctx).Err()
	if err != nil {
		return err
	}

	return nil
}

func FetchPost(id string) (*http.Response, error) {
	resp, err := http.Get(config.POST_URL + "/" + id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func FetchTodo(id string) (*http.Response, error) {
	resp, err := http.Get(config.TODO_URL + "/" + id)
	if err != nil {
		return nil, err
	}
	return resp, err
}
