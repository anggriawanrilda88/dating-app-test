package entities

import (
	"errors"
	"os"
	"time"

	vo_user "github.com/dating-app-test/src/domain/value_object/user"
	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseEntity
	id          uint64
	email       string
	password    string
	status      vo_user.UserStatus
	deletedAt   *time.Time
	accessToken string
}

func (v *User) GetID() uint64               { return v.id }
func (v *User) GetEmail() string            { return v.email }
func (v *User) GetPassword() string         { return v.password }
func (v *User) SetPassword(password string) { v.password = password }

func (v *User) GetHashPassword() (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(v.password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}

func (v *User) GetHashVerifyPassword(hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(v.password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return errors.New("incorrect password")
	}

	return nil
}

func (v *User) GetAccessToken() string {
	return v.accessToken
}

func (v *User) SetAccessToken() error {
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["exp"] = time.Now().Add(time.Minute * 43200).Unix() // 1 day expired
	atClaims["id"] = v.id
	atClaims["email"] = v.email
	atClaims["status"] = v.status
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return err
	}

	v.accessToken = accessToken

	return nil
}

func (v *User) GetStatus() vo_user.UserStatus { return v.status }
func (v *User) GetStatusString() string       { return string(v.status) }
func (v *User) GetDeletedAt() *time.Time      { return v.deletedAt }
func (v *User) MarkAsDeleted() {
	timeNow := time.Now()
	v.deletedAt = &timeNow
	v.markAsDeleted = true
}

func (v *User) RestoreFromDeleted() {
	v.markAsDeleted = false
}

// function to handle when new entity created all bussiness logic applied here
func CreateUser(
	email string,
	password string,
	status vo_user.UserStatus,
) (*User, error) {
	user := &User{
		BaseEntity: BaseEntity{
			isNew:     true,
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
		email:    email,
		password: password,
		status:   status,
	}

	// validate relate to bussiness logic
	if err := validation.ValidateStruct(
		user,
		validation.Field(&user.email, validation.Required, is.Email),
		validation.Field(&user.password, validation.Required, validation.Length(6, 100)),
	); err != nil {
		return nil, err
	}

	hashPassword, err := user.GetHashPassword()
	if err != nil {
		return nil, err
	}

	user.SetPassword(hashPassword)

	return user, nil
}

// function to handle when entity has update
func MakeUser(
	id uint64,
	email string,
	password string,
	status vo_user.UserStatus,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) (*User, error) {
	return &User{
		BaseEntity: BaseEntity{
			isModified: true,
			createdAt:  createdAt,
			updatedAt:  updatedAt,
		},
		id:        id,
		email:     email,
		password:  password,
		status:    status,
		deletedAt: deletedAt,
	}, nil
}
