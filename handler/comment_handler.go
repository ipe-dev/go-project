package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/go-project/packages/comment"
)

func CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		e := comment.Create(c)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}

func GetComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		com, e := comment.Get(c)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, com)
		}
	}
}
