package mobileapp_route

import (
	"net/http"

	"github.com/dating-app-test/src/infra/auth/jwt"
	"github.com/dating-app-test/src/interface/rest/middleware"
	"github.com/dating-app-test/src/interface/rest/v1/mobile_app/handlers"

	"github.com/gin-gonic/gin"
)

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "check_success",
	})
}

func UserRoutes(r *gin.RouterGroup, tk *jwt.Token, handler handlers.UserHandler) {
	r.POST("/", handler.RegistrationUser)
	r.POST("/login", handler.LoginUser)
	r.GET("/:id", middleware.AuthMiddleware(), handler.GetUser)
}
