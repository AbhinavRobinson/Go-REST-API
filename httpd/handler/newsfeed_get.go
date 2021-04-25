package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsfeedGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "found me",
		})
	}
}
