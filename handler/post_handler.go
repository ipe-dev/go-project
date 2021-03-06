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
		p, e := post.Get(c)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, p)
		}
	}
}
func UpdatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		e := post.Update(c)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}
func DeletePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		e := post.Delete(c)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}
func ListPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		p, e := post.List(c)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, p)
		}
	}
}
