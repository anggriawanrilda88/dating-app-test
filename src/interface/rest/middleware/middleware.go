package middleware

import (
	"net/http"

	"github.com/dating-app-test/src/infra/auth/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(routeAlias string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsMap, err := jwt.TokenValid(c.Request, routeAlias)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  err.Error(),
			})
			c.Abort()
			return
		}

		// assign userId to param
		c.Params = append(c.Params, gin.Param{
			Key:   "userId",
			Value: claimsMap["user_id"].(string),
		})

		// assign user role code to param
		c.Params = append(c.Params, gin.Param{
			Key:   "email",
			Value: claimsMap["email"].(string),
		})

		// assign user name
		c.Params = append(c.Params, gin.Param{
			Key:   "status",
			Value: claimsMap["status"].(string),
		})

		c.Next()
	}
}
