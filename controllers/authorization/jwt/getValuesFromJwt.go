package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

func GetValueFromJWT(token string, keys []string) (map[string]any, error) {
	values := make(map[string]any)
	if parsedToken, err := ParseToken(token); err == nil {
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			for _, key := range keys {
				values[key] = claims[key]
			}
		}
	}
	return values, nil
}
