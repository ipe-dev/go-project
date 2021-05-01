package comment

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/go-project/database"
)

type Comment struct {
	ID         int    `json:"id"`
	PostId     int    `json:"post_id"`
	UserId     int    `json:"user_id"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

func Create(c *gin.Context) (err error) {
	var comment Comment
	Db := database.Db
	tx := Db.Begin()
	err = c.BindJSON(&comment)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}
	comment.CreateDate = time.Now().Format("2006/01/02 15:05:05")
	fmt.Println(time.Now())
	err = Db.Create(&comment).Error
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}
	tx.Commit()
	return

}

func Update(c *gin.Context) (err error) {
	var comment Comment
	Db := database.Db
	err = c.BindJSON(&comment)
	if err != nil {
		log.Println(err)
		return
	}
	err = Db.Model(&comment).Update(comment).Error
	if err != nil {
		log.Println(err)
		return
	}
	return

}

func Get(c *gin.Context) (comment Comment, err error) {
	var request GetCommentRequest
	err = c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		return
	}
	Db := database.Db
	err = Db.Table("comments").Where("id = $1", request.Id).First(&comment).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func Delete(c *gin.Context) (err error) {
	var comment Comment
	err = c.BindJSON(&comment)
	if err != nil {
		log.Println(err)
		return
	}
	Db := database.Db
	err = Db.Delete(comment).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}
