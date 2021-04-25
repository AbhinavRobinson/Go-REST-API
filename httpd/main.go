package main

import (
	"Go-REST-API/httpd/handler"

	"Go-REST-API/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()

	// test route /ping -> pong
	r.GET("/ping", handler.PingGet())

	// newsfeed GET route
	r.GET("/newsfeed", handler.NewsfeedGet(feed))

	// newsfeed POST route
	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	// boot up router
	r.Run()
}
