package rest

import (
	"net/http"

	"github.com/dating-app-test/src/infra/auth/jwt"
	"github.com/dating-app-test/src/infra/helpers"
	"github.com/dating-app-test/src/infra/persistence/postgresql"
	handler_mb_app_v1 "github.com/dating-app-test/src/interface/rest/v1/mobile_app/handlers"
	mobileapp_route "github.com/dating-app-test/src/interface/rest/v1/mobile_app/routes"
	"github.com/gin-gonic/gin"
)

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "check_success",
	})
}

func NewRoutes(r *gin.Engine, token *jwt.Token, services *postgresql.Repositories, helpers *helpers.Helpers) {
	// api check health
	r.GET("/health", CheckHealth)

	MobileAppRouteV1(r, token, services, helpers)
}

func MobileAppRouteV1(r *gin.Engine, token *jwt.Token, services *postgresql.Repositories, helpers *helpers.Helpers) {
	// route v1 group
	api := r.Group("/dating-app-test/api/v1")

	// handlers
	userHandler := handler_mb_app_v1.NewUserHandler(services.Users, token, helpers)

	// users route
	mobileapp_route.UserRoutes(api.Group("/users"), token, userHandler)
}
