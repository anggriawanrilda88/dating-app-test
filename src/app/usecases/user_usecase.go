package usecases

import (
	"context"

	"github.com/dating-app-test/src/app/dto"
	"github.com/dating-app-test/src/domain/entities"
	"github.com/dating-app-test/src/domain/repositories"
	vo_user "github.com/dating-app-test/src/domain/value_object/user"
	"github.com/dating-app-test/src/infra/helpers"
)

type userApp struct {
	us repositories.UserRepository
	h  *helpers.Helpers
}

type UserAppInterface interface {
	RegistrationUser(ctx context.Context, dto *dto.RegistrationUserDTO) error
}

func NewUsers(us repositories.UserRepository, h *helpers.Helpers) UserAppInterface {
	return &userApp{
		us: us,
		h:  h,
	}
}

func (u *userApp) RegistrationUser(ctx context.Context, dto *dto.RegistrationUserDTO) error {
	newUser, err := entities.CreateUser(
		dto.Email,
		dto.Password,
		vo_user.UserStatusActive,
	)
	if err != nil {
		return err
	}

	return u.us.SaveUser(newUser)
}
