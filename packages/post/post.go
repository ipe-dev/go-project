package post

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/go-project/database"
	"github.com/ipe-dev/go-project/packages/comment"
	"github.com/ipe-dev/go-project/packages/user"
)

type Post struct {
	ID         int               `json:"id"`
	Title      string            `json:"title"`
	UserID     int               `json:"user_id"`
	Content    string            `json:"content"`
	Comments   []comment.Comment `json:"comments"`
	CreateDate string            `json:"create_date"`
	User       user.User         `json:"user"`
}

func Create(c *gin.Context) (err error) {
	var post Post
	err = c.BindJSON(&post)
	if err != nil {
		log.Println(err)
		return err
	}
	Db := database.Db
	post.CreateDate = time.Now().Format("2006/01/02 15:05:05")
	err = Db.Create(&post).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func Update(c *gin.Context) (err error) {
	var post Post
	err = c.BindJSON(&post)
	if err != nil {
		log.Println(err)
		return
	}
	Db := database.Db
	err = Db.Model(&post).Updates(post).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func Get(c *gin.Context) (post Post, err error) {
	var request GetRequest
	err = c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		return
	}
	Db := database.Db

	err = Db.Where("id = ?", request.ID).Preload("User").Preload("Comments").Preload("Comments.User").First(&post).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func Delete(c *gin.Context) (err error) {
	var post Post
	err = c.BindJSON(&post)
	if err != nil {
		log.Println(err)
		return
	}
	Db := database.Db
	err = Db.Delete(post).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func List(c *gin.Context) (posts []Post, err error) {
	var request ListRequest
	err = c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		return
	}
	Db := database.Db
	Db = Db.Table("posts")
	if request.Word != "" {
		Db.Where("title LIKE ? OR content LIKE ?", "%"+request.Word+"%", "%"+request.Word+"%")
	}
	if request.TagID != 0 {
		Db.Where("tag_id = ?", request.TagID)
	}
	err = Db.Preload("User").Preload("Comments").Preload("Comments.User").Find(&posts).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}
