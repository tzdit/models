package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//JWTClaim define auth claim struct
type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

//LoginRequest defines logs struct
type LoginRequest struct {
	Password string `json:"password" form:"password" validate:"required,min=1"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Csrf     string `json:"csrf" form:"csrf" validate:"required"`
}

//Defines UserStructure
type User struct {
	ID             int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name           string    `json:"name" form:"name" validate:"required"`
	Email          string    `json:"email" form:"email" validate:"required,email"`
	Password       string    `json:"password" form:"password"  validate:"required,min=1"`
	Active         bool      `json:"active" form:"active" validate:"omitempty"`
	Token          string    `json:"token,omitempty"`
	UserCategoryID int32     `json:"user_category_id,omitempty"`
	LoginAttempt   int       ` json:"login_attempt,omitempty"`
	VerifiedAt     time.Time `json:"verified_at,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

//Defines UserStructure
type UserEdit struct {
	ID    int32  `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name  string `json:"name" form:"name" validate:"required"`
	Email string `json:"email" form:"email" validate:"required,email"`
}

type UserID struct {
	ID int32 `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
}

type UserRole struct {
	ID    int32   `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Roles []int32 `json:"roles" form:"roles[]" validate:"required"`
}
