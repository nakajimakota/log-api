package main

import (
	// for debuging
	"fmt"
	// "reflect"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
	

	"app/db"
	"app/handlers"
	"app/models"
	"app/lib"
)

func setupRouter() *gin.Engine {

	fmt.Println("OKOKOK")

	r := gin.Default()
	// r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8010",
		},
		AllowMethods: []string{
			"POST",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"X-CSRFToken",
		},
		AllowCredentials: true,
		MaxAge: 24 * time.Hour,
	  }))
	// log
	log := r.Group("/log")
	log.Use(lib.AuthHeader())
	{
		log.GET("/list/:userId", handlers.GetLogList)
		log.POST("/create/:userId", handlers.CreateLog)
		log.POST("/delete/:logId", handlers.DeleteLog)
	}
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
