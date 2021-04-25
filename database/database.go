package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var Db *gorm.DB

func Connect() {
	var err error
	Db, err = gorm.Open("postgres", "user=root dbname=go_project password=go_project sslmode=disable")
	if err != nil {
		panic(err)
	}
}
