package requests

import (
	"errors"

	"github.com/dating-app-test/src/app/dto"
	"github.com/gin-gonic/gin"
)

// IOrderRequest ...
type UserRequest interface {
	Validate(c *gin.Context) (*dto.RegistrationUserDTO, error)
}

type RegistrationUser struct {
	Email    string
	Password string
}

func (req *RegistrationUser) Validate(c *gin.Context) (*dto.RegistrationUserDTO, error) {
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return nil, errors.New("error request body")
	}

	return &dto.RegistrationUserDTO{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}
