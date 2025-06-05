package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Token- "checkmarx"
		c.Writer.Header().Set("X-Info", "Add token - Checkmarx")
		token := c.GetHeader("Authorization")
		if token != "Bearer checkmarx" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized or missing token goto Authorization ,On swagger you would see 'Authorization'. insert access Token -- token is 'Bearer checkmarx'",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
