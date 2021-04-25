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
	if err := c.BindJSON(&post); err != nil {
		log.Println(err)
		return err
	}
	Db := database.Db
	err := Db.Create(&post).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func Update() {

}

func Get() {

}

func Delete() {

}
