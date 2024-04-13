package model

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/miladibra10/vjson"
	"github.com/sarastee/avito-test-assignment/internal/utils/validator"
)

// CreateUser model struct
type CreateUser struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

// AuthUser model struct
type AuthUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UserClaims model struct
type UserClaims struct {
	jwt.RegisteredClaims
	UserID   int64  `json:"id"`
	UserName string `json:"username"`
	Role     string `json:"role"`
}

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password string `json:"password_hash"`
}

// Token model struct
type Token struct {
	Token string `json:"token"`
}

func ValidateCreateUser(data []byte) error {
	log.Printf(string(data))
	schema := validator.NewSchema(
		vjson.String("name").Required(),
		vjson.String("role").Choices("ADMIN", "USER").Required(),
		vjson.String("password").Required(),
	)

	return schema.ValidateBytes(data)
}

func ValidateAuthUser(data []byte) error {
	log.Printf(string(data))
	schema := validator.NewSchema(
		vjson.String("name").Required(),
		vjson.String("password").Required(),
	)

	return schema.ValidateBytes(data)
}
