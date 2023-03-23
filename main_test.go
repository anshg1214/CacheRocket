package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anshg1214/CacheRocket/config"
	"github.com/anshg1214/CacheRocket/controller"
	"github.com/anshg1214/CacheRocket/utils"
)

var router *gin.Engine

// TestMain is the entry point for all tests
func TestMain(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		status      int
		cacheStatus bool
	}{
		{
			name:   "Test Get Post",
			path:   "/posts/1",
			status: http.StatusOK,
		},
		{
			name:   "Test Get Post 404",
			path:   "/posts/10000",
			status: http.StatusNotFound,
		},
		{
			name:   "Test Get Todo",
			path:   "/todos/1",
			status: http.StatusOK,
		},
		{
			name:   "Test Get Todo 404",
			path:   "/todos/10000",
			status: http.StatusNotFound,
		},
	}

	gin.SetMode(gin.TestMode)

	// Setup redis
	err := config.PingRedis()
	assert.NoError(t, err, "❌ Redis is not running")

	// Flush redis
	err = utils.FlushCache()
	assert.NoError(t, err, "Error flushing cache: %v", err)

	router = gin.Default()

	// Routes
	router.GET("/", controller.GetHome)
	router.GET("/posts/:id", controller.GetPost)
	router.GET("/todos/:id", controller.GetTodo)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testGet(t, test.path, test.status, test.cacheStatus)
		})
	}
}

func testGet(t *testing.T, path string, status int, cacheStatus bool) {

	// Create a request to send to the above route
	req, err := http.NewRequest("GET", path, nil)
	assert.NoError(t, err, "Error creating request: %v", err)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Send the request to the router
	router.ServeHTTP(w, req)

	// Check the status code is what we expect
	assert.Equal(t, status, w.Code, "handler returned wrong status code: got %v want %v", w.Code, status)

	// Check the response body is what we expect
	body, err := io.ReadAll(w.Body)
	assert.NoError(t, err, "Error reading response body: %v", err)

	checkValidJson := json.Valid(body)
	assert.Equal(t, true, checkValidJson)
}
