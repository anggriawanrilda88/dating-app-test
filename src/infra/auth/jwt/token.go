package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dating-app-test/src/domain/entities"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type Token struct{}

func NewToken() *Token {
	return &Token{}
}

type TokenInterface interface {
	CreateToken(user *entities.User, roleCode *string) (*TokenDetails, error)
	ExtractTokenMetadata(*http.Request) (*AccessDetails, error)
}

// Token implements the TokenInterface
var _ TokenInterface = &Token{}

func (t *Token) CreateToken(user *entities.User, roleCode *string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 43200).Unix()
	td.TokenUuid = uuid.NewV4().String()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.TokenUuid
	atClaims["user_id"] = user.GetID()
	atClaims["email"] = user.GetEmail()
	atClaims["status"] = user.GetStatus()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func TokenValid(r *http.Request) (jwt.MapClaims, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.Claims)
	if !ok && !token.Valid {
		return nil, err
	}

	claimsMap := claims.(jwt.MapClaims)
	return claimsMap, nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)

	if tokenString == "" {
		return nil, errors.New("forbidden access")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// get the token from the request body
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (t *Token) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			TokenUuid: accessUuid,
			UserId:    userId,
		}, nil
	}
	return nil, err
}
