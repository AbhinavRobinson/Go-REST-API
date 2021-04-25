package handler

import (
	"Go-REST-API/platform/newsfeed"

	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsfeedPost(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "found me",
		})
	}
}
