package helpers

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/shiv122/go-todo/config"
)

type JwtHelper struct{}

func (jh *JwtHelper) ExtractClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		return []byte(config.App.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("error in jwt claims")
	}

	return claims, nil
}
