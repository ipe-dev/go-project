package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/go-project/packages/post"
)

func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
		e := post.Create(c.Request)
		c.JSON(http.StatusOK, gin.H{
			"hoge": "fuga",
		})
	}
}

func GetPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hoge": "fuga",
		})
	}
}
