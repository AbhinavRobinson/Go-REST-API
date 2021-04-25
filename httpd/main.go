package main

import (
	"Go-REST-API/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// test route /ping -> pong
	r.GET("/ping", handler.PingGet())

	// newsfeed GET route
	r.GET("/newsfeed", handler.NewsfeedGet())

	// newsfeed POST route
	r.POST("/newsfeed", handler.NewsfeedGet())

	// boot up router
	r.Run()
}
