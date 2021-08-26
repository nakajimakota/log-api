package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	// DEBUG:
	// "log"

	"github.com/gin-gonic/gin"

	"app/controllers"
	// "app/lib"
	// "app/models"
)

func CreateLog(c *gin.Context) {
	// var jsonData models.Log
	userId, _ := strconv.Atoi(c.Param("userId"))
	fmt.Println(userId)
	// fmt.Println(c.BindJSON(&jsonData))
	fmt.Println(c.PostForm("title"))
	ctrl := controllers.NewLog()
	logs := ctrl.CreateLog(userId)


	fmt.Println("logs", logs)

	fmt.Println("http.StatusOK", http.StatusOK)

	c.JSON(http.StatusOK, logs)
}

func GetLogList(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))
	ctrl := controllers.NewLog()
	logs := ctrl.GetLogList(userId)

	c.JSON(http.StatusOK, logs)
}