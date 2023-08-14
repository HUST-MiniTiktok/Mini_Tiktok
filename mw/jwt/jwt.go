package jwt

import (
	"errors"

	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	ID int64
	jwt.RegisteredClaims
}

var Jwt *JWT

func init() {
	Jwt = &JWT{
		SigningKey: []byte(conf.GetConf().GetString("jwt.signingkey")),
	}
}

func GetJwt() *JWT {
	return Jwt
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token is invalid")
}