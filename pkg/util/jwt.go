package util

import (
	"fmt"
	"time"

	j "github.com/golang-jwt/jwt/v5"
	"github.com/mizmorr/rest-example/pkg/logger"
)

type jwt struct {
	secretKey []byte
}

func (jwt *jwt) _init(sk []byte) {

	jwt.secretKey = sk

}

func NewJWT(key string) *jwt {

	jwt := jwt{}
	jwt._init([]byte(key))

	return &jwt
}

var (
	l = logger.Get()
)

func (jwt *jwt) CreateToken(username, role string) (string, error) {

	claims := j.NewWithClaims(
		j.SigningMethodHS256,
		j.MapClaims{
			"sub": username,
			"iss": "rest-app",
			"aud": role,
			"exp": time.Now().Add(time.Hour).Unix(),
			"iat": time.Now().Unix(),
		})
	tokenString, err := claims.SignedString(jwt.secretKey)

	if err != nil {
		return "", err
	}

	l.Debug().Msg(fmt.Sprintf("Token claims added: %+v\n", claims))

	return tokenString, nil
}
