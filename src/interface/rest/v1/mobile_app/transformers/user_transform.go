package transformers

import (
	"github.com/dating-app-test/src/domain/entities"
)

type loginTransform struct {
	ID     uint64 `json:"id"`
	Email  string `json:"email"`
	Status string `json:"status"`
	Token  string `json:"token"`
}

// ResponseListHandler, response format for list data
func LoginTransform(user *entities.User) *loginTransform {
	return &loginTransform{
		ID:     user.GetID(),
		Email:  user.GetEmail(),
		Status: string(user.GetStatusString()),
		Token:  user.GetAccessToken(),
	}
}

type getUserTransform struct {
	ID        uint64 `json:"id"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// ResponseListHandler, response format for list data
func GetUserTransform(user *entities.User) *getUserTransform {
	return &getUserTransform{
		ID:        user.GetID(),
		Email:     user.GetEmail(),
		Status:    string(user.GetStatusString()),
		CreatedAt: user.GetCreatedAtAsISOString(),
		UpdatedAt: user.GetUpdatedAtISOString(),
	}
}
