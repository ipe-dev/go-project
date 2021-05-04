package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/go-project/packages/user"
)

// CreateUser user作成
func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		e := user.Create(c)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}

// GetUser user取得
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, e := user.Get(c)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, u)
		}
	}
}

// UpdateUser user更新
func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		e := user.Update(c)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}

// DeleteUser user削除
func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		e := user.Delete(c)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}

// ListUser user一覧取得
func ListUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, e := user.List(c)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		} else {
			c.JSON(http.StatusOK, u)
		}
	}
}
