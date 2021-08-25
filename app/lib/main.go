package lib

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    "Not Found",
		"message": "作成エラー",
	})
}