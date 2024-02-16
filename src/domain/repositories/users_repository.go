package repositories

import (
	"github.com/dating-app-test/src/domain/entities"
)

type UserRepository interface {
	SaveUser(*entities.User) error
	UpdateUser(*entities.User) error
	GetUser(string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	GetUserByEmailAndPassword(*entities.User) (*entities.User, *string, error)
}
