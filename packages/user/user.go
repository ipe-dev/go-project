package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/go-project/database"
)

type User struct {
	ID         int    `json:"id" gorm:"primary_key"`
	LoginID    string `json:"login_id" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Detail     string `json:"detail"`
	LegendID   int    `json:"legend_id" binding:"required"`
	RankID     int    `json:"rank_id" binding:"required"`
	CreateDate string `json:"create_date"`
}

func Create(c *gin.Context) (err error) {
	var user User
	err = c.BindJSON(&user)
	if err != nil {
		log.Println(err)
		return
	}

	Db := database.Db
	err = Db.Create(&user).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}
func Get(c *gin.Context) (user User, err error) {
	var request GetUserReqest
	err = c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		return
	}
	Db := database.Db
	err = Db.Where("id = $1", request.ID).First(&user).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}
func Update(c *gin.Context) (err error) {
	var user User
	err = c.BindJSON(&user)
	if err != nil {
		log.Println(err)
		return
	}
	Db := database.Db
	err = Db.Model(&user).Updates(user).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}
func Delete(c *gin.Context) (err error) {
	var user User
	err = c.BindJSON(&user)
	if err != nil {
		log.Println(err)
		return
	}
	Db := database.Db
	err = Db.Delete(user).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}
func List(c *gin.Context) (users []User, err error) {
	var request GetListUserReqest
	err = c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		return
	}
	Db := database.Db
	Db.Table("users")
	if request.LoginID != "" {
		Db = Db.Where("login_id LIKE ?", "%"+request.LoginID+"%")
	}
	if request.Word != "" {
		Db = Db.Where("name LIKE ? OR detail LIKE ?", "%"+request.Word+"%", "%"+request.Word+"%")
	}
	if request.RankID != 0 {
		Db = Db.Where("rank_id", request.RankID)
	}
	Db.Debug().Find(&users)

	return
}
