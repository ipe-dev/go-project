package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/go-project/handler"
)

func main() {
	r := gin.Default()
	r.POST("/api/post/create", handler.CreatePost())
	r.Run()
}
