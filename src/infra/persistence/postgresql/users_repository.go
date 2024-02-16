package postgresql

import (
	"errors"

	"github.com/dating-app-test/src/domain/entities"
	"github.com/dating-app-test/src/domain/repositories"
	"github.com/dating-app-test/src/infra/models"

	"golang.org/x/crypto/bcrypt"
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

func (r *UsersRepo) UpdateUser(user *entities.User) error {
	tx := r.db.Begin()
	err := tx.Model(&user).Updates(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *UsersRepo) GetUser(id string) (*entities.User, error) {
	var user entities.User
	err := r.db.Joins("RoleAssignments").
		Joins("Stations").Where(`"users"."id" = ?`, id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UsersRepo) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where(`"users"."email" = ?`, email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UsersRepo) GetUserByEmailAndPassword(u *entities.User) (*entities.User, *string, error) {
	var user entities.User
	err := r.db.Where("email = ?", u.GetEmail()).Take(&user).Error
	if err != nil {
		return nil, nil, errors.New("database error")
	}
	//Verify the password
	err = user.GetHashVerifyPassword(u.GetPassword())
	if err != nil {
		return nil, nil, err
	}
	// get role
	var roleCode string
	err = r.db.Raw(`
	select rl.code from role_assignments as ra 
	join roles as rl on ra.role_id = rl.id where user_id = $1`, user.GetID()).Scan(&roleCode).Error

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, nil, err
	}

	return &user, &roleCode, nil
}
