package models

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}
