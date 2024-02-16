package postgresql

import (
	"github.com/dating-app-test/src/domain/entities"
	"github.com/dating-app-test/src/domain/repositories"
	"github.com/dating-app-test/src/infra/models"

	"gorm.io/gorm"
)

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *UsersRepo {
	return &UsersRepo{db}
}

// UsersRepo implements the repository.UserRepository interface
var _ repositories.UserRepository = &UsersRepo{}

func (r *UsersRepo) SaveUser(user *entities.User) error {
	userModel := models.ToUserModel(user)
	tx := r.db.Begin()
	err := tx.Create(&userModel).Error
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (r *UsersRepo) GetUser(id uint64) (*entities.User, error) {
	var model models.User
	err := r.db.Where(`"users"."id" = ?`, id).Take(&model).Error
	if err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}

func (r *UsersRepo) GetUserByEmail(email string) (*entities.User, error) {
	var model models.User
	err := r.db.Where(`"users"."email" = ?`, email).Take(&model).Error
	if err != nil {
		return nil, err
	}

	return model.ToEntity(), nil
}
