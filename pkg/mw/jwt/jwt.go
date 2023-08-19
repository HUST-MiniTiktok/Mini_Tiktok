package jwt

import (
	"errors"

	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	SigningKey []byte
}

type UserClaims struct {
	ID int64
	jwt.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(conf.GetConf().GetString("jwt.signingkey")),
	}
}

func (j *JWT) CreateToken(claims UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ExtractClaims(tokenString string) (*UserClaims, error) {
	if tokenString == "" {
		return nil, errors.New("empty token")
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
