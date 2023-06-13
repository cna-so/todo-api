package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

func GetValueFromJWT(token *jwt.Token, keys []string) (map[string]any, error) {
	values := make(map[string]any)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for _, key := range keys {
			values[key] = claims[key]
		}
	}
	return values, nil
}
