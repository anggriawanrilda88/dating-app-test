package requests

import (
	"errors"
	"strconv"

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

type LoginUser struct {
	Email    string
	Password string
}

func (req *LoginUser) Validate(c *gin.Context) (*dto.LoginUserDTO, error) {
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return nil, errors.New("error request body")
	}

	return &dto.LoginUserDTO{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}

type GetUser struct {
	ID uint64
}

func (req *GetUser) Validate(c *gin.Context) (*dto.GetUserDTO, error) {
	// req validate, convert req string id to uint64
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return nil, errors.New("error request id")
	}
	req.ID = id

	return &dto.GetUserDTO{
		ID: req.ID,
	}, nil
}
