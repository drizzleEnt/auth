package model

import (
	"github.com/dgrijalva/jwt-go"
)

const (
	ExamplePath = "/user/v2/create"
)

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}
