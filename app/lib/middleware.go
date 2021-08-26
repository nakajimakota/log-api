package lib

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestHeader := c.GetHeader("Content-Type")
		fmt.Println("AS", requestHeader)
		if strings.Index(requestHeader, "application/json") != 0 {
			c.JSON(http.StatusUnsupportedMediaType,
				gin.H{
					"code":    "Unsupported Media Type",
					"message": "application/jsonでリクエストして下さい",
				})
			c.Abort()
		}
		c.Next()
	}
}
