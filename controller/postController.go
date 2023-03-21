package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/anshg1214/CacheRocket/config"
	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	// Get the post ID from the URL
	postID := c.Param("id")

	// Get the post from the URL JSONPlaceholder API
	response, err := http.Get(config.POST_URL + "/" + postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Close the response body
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if the response body is valid JSON
	checkValidJson := json.Valid(body)

	if !checkValidJson {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON"})
		return
	}

	// Send the response body to the client
	c.Data(http.StatusOK, "application/json", body)
}
