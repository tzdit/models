package models

import (
	"github.com/dgrijalva/jwt-go"
)

// JwtClaims struct
type JwtClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// Login struct
type Login struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

/*
// AuthLogin Login Query
func AuthLogin(email string, password string) *User {
	user := new(User)
	res := database.DB.Where("email = ? and password = ?", email, password).First(&user)
	if res.Error == nil {

		json, err := json.Marshal(&user)
		if err == nil {
			database.DB.Set(  user.ID, json)

		}

		return user
	}
	return nil
}*/
