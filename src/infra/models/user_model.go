package models

import (
	"time"

	"github.com/dating-app-test/src/domain/entities"
	vo_user "github.com/dating-app-test/src/domain/value_object/user"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `json:"id" gorm:"column:id;primaryKey"`
	Email     string         `json:"email" gorm:"column:email"`
	Password  string         `json:"password" gorm:"column:password"`
	Status    string         `json:"status" gorm:"column:status"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at"`
}

func (m User) ToEntity() *entities.User {
	deletedAt := &m.DeletedAt.Time
	entity, _ := entities.MakeUser(
		m.ID,
		m.Email,
		m.Password,
		vo_user.UserStatus(m.Status),
		m.CreatedAt,
		m.UpdatedAt,
		deletedAt,
	)

	return entity
}

func ToUserModel(data *entities.User) *User {
	return &User{
		ID:        data.GetID(),
		Email:     data.GetEmail(),
		Password:  data.GetPassword(),
		Status:    string(data.GetStatus()),
		CreatedAt: data.GetCreatedAt(),
		UpdatedAt: data.GetUpdatedAt(),
	}
}
