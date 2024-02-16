package handlers

import (
	"github.com/dating-app-test/src/app/usecases"
	"github.com/dating-app-test/src/domain/repositories"
	"github.com/dating-app-test/src/infra/auth/jwt"
	"github.com/dating-app-test/src/infra/helpers"
	"github.com/dating-app-test/src/interface/rest/response"
	"github.com/dating-app-test/src/interface/rest/v1/mobile_app/requests"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	RegistrationUser(c *gin.Context)
}

type userHandler struct {
	userService repositories.UserRepository
	token       jwt.TokenInterface
	helpers     *helpers.Helpers
	usecase     usecases.UserAppInterface
}

func NewUserHandler(userService repositories.UserRepository, token jwt.TokenInterface, helpers *helpers.Helpers) UserHandler {
	usecase := usecases.NewUsers(userService, helpers)
	return &userHandler{
		userService: userService,
		token:       token,
		helpers:     helpers,
		usecase:     usecase,
	}
}

// Create new user
func (h *userHandler) RegistrationUser(c *gin.Context) {
	req := requests.RegistrationUser{}
	dto, err := req.Validate(c)
	if err != nil {
		response.ErrorHandler(c, err)
		return
	}

	err = h.usecase.RegistrationUser(c.Request.Context(), dto)
	if err != nil {
		response.ErrorHandler(c, err)
		return
	}

	response.ResponseJSON(c, "User successfully registered", nil)
}
