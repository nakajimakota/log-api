package main

import (
	// for debuging
	// "fmt"
	// "reflect"
	"github.com/gin-gonic/gin"

	"app/db"
	// "app/handlers"
	"app/models"
)

func setupRouter() *gin.Engine {

	r := gin.Default()
	// r.Use(cors.Default())
	// log
	// tenant := r.Group("/log")
	// tenant.Use(common.AuthHeader())

	return r
}

var router = setupRouter()

func main() {
	mysql := db.Init()
	mysql.AutoMigrate(&models.Log{})
	defer mysql.Close()

	r := router
	r.Run(":8011")
}
