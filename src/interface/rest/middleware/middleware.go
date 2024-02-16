package middleware

import (
	"net/http"

	"github.com/dating-app-test/src/infra/auth/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsMap, err := jwt.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  err.Error(),
			})
			c.Abort()
			return
		}

		// set params to context
		c.Set("userID", claimsMap["id"])
		c.Set("email", claimsMap["email"])
		c.Set("status", claimsMap["status"])

		c.Next()
	}
}
