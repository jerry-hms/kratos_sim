package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	UserId    int64  `json:"user_id"`
	CreatedAt string `json:"created_at"`
}
