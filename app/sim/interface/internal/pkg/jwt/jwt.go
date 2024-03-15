package jwt

import (
	"context"
	"encoding/json"
	"errors"
	ja "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/golang-jwt/jwt/v4"
)

func NewJwtSign(key string) *JwtSign {
	return &JwtSign{
		SignKey: []byte(key),
	}
}

func GetAuth(ctx context.Context) (*Claims, error) {
	var myClaims *Claims
	claims, ok := ja.FromContext(ctx)
	if !ok {
		return nil, errors.New("数据解析失败")
	}
	claimsMap, ok := claims.(*jwt.MapClaims)
	if !ok {
		return nil, errors.New("数据解析失败")
	}
	bytes, _ := json.Marshal(claimsMap)
	err := json.Unmarshal(bytes, &myClaims)
	if err != nil {
		return nil, err
	}
	return myClaims, nil
}

type JwtSign struct {
	SignKey []byte
}

// 创建token
func (j *JwtSign) CreateToken(c Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	signToken, err := token.SignedString(j.SignKey)
	if err != nil {
		return "", err
	}
	return signToken, nil
}

// ParseToken 解析token
func (j *JwtSign) ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SignKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
