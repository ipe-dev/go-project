package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/go-project/packages/post"
)

func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		e := post.Create(c)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}

func GetPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hoge": "fuga",
		})
	}
}
