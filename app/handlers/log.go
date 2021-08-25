package handlers

import (
	"net/http"
	// "strconv"

	// DEBUG:
	// "log"

	"github.com/gin-gonic/gin"

	"app/controllers"
	"app/lib"
	// "app/models"
)

func CreateLog(c *gin.Context) {
	// params := c.Request.URL.Query()
	// tenantId, err := strconv.Atoi(c.Param("tenantId"))
	// if err != nil {
	// 	BadRequestMessage(c, "テナントIDが整数でありません")
	// 	return
	// }
	ctrl := controllers.NewLog()
	is_create := ctrl.CreateLog()
	if is_create == false {
		lib.NotFound(c)
	}
	c.JSON(http.StatusOK, true)
}