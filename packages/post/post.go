package post

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/go-project/database"
	"github.com/ipe-dev/go-project/packages/comment"
)

type Post struct {
	ID       int               `json:"id"`
	Title    string            `json:"title"`
	UserName string            `json:"user_name"`
	Content  string            `json:"content"`
	Comment  []comment.Comment `json:"comments"`
}

func Create(c *gin.Context) error {
	var post Post
	var err error
	err = c.BindJSON(&post)
	if err != nil {
		log.Println(err)
		return err
	}
	Db := database.Db
	err = Db.Create(&post).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func Update(c *gin.Context) error {
	var post Post
	var err error
	err = c.BindJSON(&post)
	if err != nil {
		log.Println(err)
		return err
	}
	Db := database.Db
	err = Db.Model(&post).Updates(post).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func Get(c *gin.Context) (Post, error) {
	var request GetRequest
	var post Post
	err := c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		return post, err
	}
	Db := database.Db
	err = Db.Where("id = $1", request.ID).First(&post).Error
	if err != nil {
		log.Println(err)
		return post, err
	}
	return post, err
}

func Delete() {

}
