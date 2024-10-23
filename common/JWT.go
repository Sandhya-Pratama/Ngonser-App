package common

import "github.com/golang-jwt/jwt/v5"

// JwtCustomClaims custom claims
type JwtCustomClaims struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Roles string `json:"roles"`
	jwt.RegisteredClaims
}
