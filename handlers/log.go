package handlers

import (
	// "fmt"
	"net/http"
	"strconv"

	// DEBUG:
	// "log"

	"github.com/gin-gonic/gin"

	"controllers"
	// "lib"
	"models"
	// "log"
)
func CreateLog(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))
	var log models.Log
	c.ShouldBindJSON(&log)
	requestData := map[string]string{
		"Title": log.Title,
		"ImageURL": log.ImageURL,
		"SiteURL": log.SiteURL,
	}

	ctrl := controllers.NewLog()
	err := ctrl.CreateLog(userId, requestData)

	c.JSON(http.StatusOK, err)
}

func GetLogList(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))
	ctrl := controllers.NewLog()
	logs := ctrl.GetLogList(userId)

	c.JSON(http.StatusOK, logs)
}

func DeleteLog(c *gin.Context) {
	logId, _ := strconv.Atoi(c.Param("logId"))
	ctrl := controllers.NewLog()
	res := ctrl.DeleteLog(logId)

	c.JSON(http.StatusOK, res)
}