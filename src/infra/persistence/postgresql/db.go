package postgresql

import (
	"fmt"
	"os"

	"github.com/dating-app-test/src/domain/repositories"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type Repositories struct {
	Users repositories.UserRepository
	db    *gorm.DB
}

func New(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// show hide debug postgresql
	debug := os.Getenv("DB_DEBUG")
	if debug == "true" {
		db = db.Debug()
	}

	return &Repositories{
		Users: NewUsersRepository(db),
		db:    db,
	}, nil
}
