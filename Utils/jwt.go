package Utils

import (
	"errors"
	"time"

	"github.com/bigby/project/Config"
	"github.com/dgrijalva/jwt-go"
)

type JWTCliams struct {
	jwt.StandardClaims
	UserID   int64
	UserType Config.UserType
	SignKey  []byte
	RdsKey   []byte
}

func GetToken(data map[string]interface{}, secret string) (string, error) {
	claims := JWTCliams{
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * time.Duration(Config.TokenExpireTime)).Unix(),
		},
		data["Id"].(int64),
		data["UserType"].(Config.UserType),
		data["SignKey"].([]byte),
		[]byte(data["RdsKey"].(string)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("服务器繁忙")
	}
	return signedToken, nil
}

func ParseToken(token string, secret string) (*JWTCliams, error) {
	secretByte := []byte(secret)
	tokenClaims, err := jwt.ParseWithClaims(token, &JWTCliams{}, func(token *jwt.Token) (interface{}, error) {
		return secretByte, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*JWTCliams); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	//
	return nil, err
}
