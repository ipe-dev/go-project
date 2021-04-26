package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/go-project/database"
	"github.com/ipe-dev/go-project/handler"
)

func init() {
	database.Connect()
}

func main() {
	r := gin.Default()
	r.POST("/api/post/create", handler.CreatePost())
	r.POST("/api/post/get", handler.GetPost())
	r.Run()
}
