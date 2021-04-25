package main

import (
	"Go-REST-API/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// test route /ping > pong
	r.GET("/ping", handler.PingGet)

	// boot up router
	r.Run()
}
