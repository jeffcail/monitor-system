package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	EXPIRE        = time.Hour * 24 * 365
	SECRET string = "xxxxxxxxxxxxxxfasdfasdfsda"
)

type JwtClaims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	RealName string `json:"real_name"` // 真实姓名
	jwt.StandardClaims
}

// GenerateToken
func GenerateToken(claims *JwtClaims) (string, error) {
	c := JwtClaims{
		claims.ID,
		claims.Username,
		claims.RealName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(EXPIRE).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(SECRET))
}

// ParseToken
func ParseToken(ts string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(ts, &JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET), nil
		})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
