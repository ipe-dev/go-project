package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/go-project/database"
	"github.com/ipe-dev/go-project/route"
)

func init() {
	database.Connect()
}

func main() {
	r := gin.Default()
	route.DefineRoutes(r)

	r.Run()
}
