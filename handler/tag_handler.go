package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/go-project/packages/tag"
)

// CreateTag tag作成
func CreateTag() gin.HandlerFunc {
	return func(c *gin.Context) {
		e := tag.Create(c)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}

// DeleteTag tag削除
func DeleteTag() gin.HandlerFunc {
	return func(c *gin.Context) {
		e := tag.Delete(c)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}

// ListTag tag一覧取得
func ListTag() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, e := tag.List(c)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, u)
		}
	}
}
