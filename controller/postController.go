package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/anshg1214/CacheRocket/config"
	"github.com/anshg1214/CacheRocket/utils"
	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	// Get the post ID from the URL
	postID := c.Param("id")

	// Get the post from the URL JSONPlaceholder API
	response, err := http.Get(config.POST_URL + "/" + postID)
	if err != nil {
		utils.ThrowInternalServerError(c, err)
		return
	}

	if response.StatusCode == http.StatusNotFound {
		utils.ThrowNotFoundError(c)
		return
	} else if response.StatusCode != http.StatusOK {
		utils.ThrowInternalServerError(c, err)
		return
	}

	// Close the response body
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		utils.ThrowInternalServerError(c, err)
		return
	}

	// Check if the response body is valid JSON
	checkValidJson := json.Valid(body)

	if !checkValidJson {
		utils.ThrowInvalidJSONError(c)
		return
	}

	c.Set("data", body)

	// Send the response body to the client
	c.Data(http.StatusOK, "application/json", body)
}
