package main

import (
	// for debuging
	"fmt"
	"os"

	// "reflect"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/nakajimakota/log-api/db"
	"github.com/nakajimakota/log-api/handlers"
	"github.com/nakajimakota/log-api/lib"
	"github.com/nakajimakota/log-api/models"
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
		MaxAge:           24 * time.Hour,
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
	os.Setenv("PORT", "8011")
	err := godotenv.Load(fmt.Sprintf("envfiles/.env", os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	port := os.Getenv("PORT")
	mysql := db.Init()
	mysql.AutoMigrate(&models.Log{})
	defer mysql.Close()
	fmt.Println(port, "WE")

	r := router
	r.Run()
	// http.ListenAndServe(":8011", r)
}
