package models

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	ID   uint   `json:"Id"`
	Role string `json:"Role"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID uint `json:"Id"`
	jwt.RegisteredClaims
}
