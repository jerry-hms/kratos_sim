package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"sync"
	"time"
)

var UserTokenIns *UserToken
var UserTokenOnce sync.Once

type UserToken struct {
	jwt *JwtSign
}

// 创建UserToken工厂
func CreateUserTokenFactory(key string) *UserToken {
	UserTokenOnce.Do(func() {
		UserTokenIns = &UserToken{
			jwt: NewJwtSign(key),
		}
	})
	return UserTokenIns
}

func (u *UserToken) GenerateToken(c Claims, expiredAt int64) (string, error) {
	c.RegisteredClaims.NotBefore = jwt.NewNumericDate(time.Now().Add(-10 * time.Second))                      // 生效开始时间 给个10秒的浮动区间
	c.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(expiredAt) * time.Second)) // 有效时间

	return u.jwt.CreateToken(c)
}

func (u *UserToken) IsNotExpired(tokenStr string) (*Claims, error) {
	claims, err := u.jwt.ParseToken(tokenStr)
	if err == nil {
		// 校验是否在有效期内
		if time.Now().Unix()-claims.ExpiresAt.Unix() < 0 {
			return claims, nil
		} else {
			return nil, errors.New("token已过期")
		}
	}
	return nil, err
}

// IsEffective 判断token是否有效
func (u *UserToken) IsEffective(tokenStr string) (*Claims, error) {
	return u.IsNotExpired(tokenStr)
}
