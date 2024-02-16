package repositories

import (
	"github.com/dating-app-test/src/domain/entities"
)

type UserRepository interface {
	SaveUser(*entities.User) error
	GetUser(id uint64) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
}
