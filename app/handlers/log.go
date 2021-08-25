package handlers

import (
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
	userId, _ := strconv.Atoi(c.Param("userId"))
	ctrl := controllers.NewLog()
	logs := ctrl.CreateLog(userId)

	c.JSON(http.StatusOK, logs)
}

func GetLogList(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))
	ctrl := controllers.NewLog()
	logs := ctrl.GetLogList(userId)

	c.JSON(http.StatusOK, logs)
}