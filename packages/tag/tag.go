package tag

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/go-project/database"
)

type Tag struct {
	ID         int    `json:"id"`
	Name       string `json:"name" binding:"required"`
	CreateDate string `json:"create_date"`
}

func Create(c *gin.Context) (err error) {
	var tag Tag
	Db := database.Db
	err = c.BindJSON(&tag)
	tx := Db.Begin()
	tag.CreateDate = time.Now().Format("2006/01/02 15:05:05")

	if err != nil {
		tx.Rollback()
		return
	}
	err = Db.Create(tag).Error
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}
	tx.Commit()
	return
}

func Delete(c *gin.Context) (err error) {
	var tag Tag
	Db := database.Db
	tx := Db.Begin()
	err = c.BindJSON(&tag)
	if err != nil {
		tx.Rollback()
		return
	}
	err = Db.Delete(tag).Error
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}
	tx.Commit()
	return
}

func List(c *gin.Context) (tags []Tag, err error) {
	var request GetListTagReqest
	err = c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		return
	}
	Db := database.Db
	Db = Db.Table("tags")
	if request.Word != "" {
		Db = Db.Where("name LIKE ?", "%"+request.Word+"%")
	}
	err = Db.Find(&tags).Error
	if err != nil {
		return
	}
	return
}
