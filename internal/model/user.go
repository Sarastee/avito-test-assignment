package model

import "github.com/golang-jwt/jwt/v5"

// UserRole object
type UserRole string

const (
	USER  UserRole = "USER"  // USER common user
	ADMIN UserRole = "ADMIN" // ADMIN admin user
)

// User model struct
type User struct {
	ID           int64    `json:"id"`
	Name         string   `json:"name"`
	Role         UserRole `json:"role"`
	PasswordHash string   `json:"password_hash"`
}

// UserClaims model struct
type UserClaims struct {
	jwt.RegisteredClaims
	UserID   string `json:"userID"`
	UserName string `json:"username"`
	Role     string `json:"role"`
}

// Token model struct
type Token struct {
	Token string `json:"token"`
}
