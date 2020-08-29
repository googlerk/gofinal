package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "November 10, 2009" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "StatusUnauthorized[401]: you don't have the permission!!"})
		// fmt.Println(http.StatusUnauthorized)
		c.Abort()
		return
	}
	c.Next()
}
