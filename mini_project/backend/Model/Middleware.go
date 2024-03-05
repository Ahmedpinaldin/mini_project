package Model

import "github.com/golang-jwt/jwt/v5"

type AuthClaimJWT struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}
