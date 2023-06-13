package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func GetValueFromJWT(token string, keys []string) (map[string]string, error) {
	values := make(map[string]string)
	if parsedToken, err := ParseToken(token); err == nil {
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			for _, key := range keys {
				values[key] = fmt.Sprintf("%s", claims[key])
			}
		}
	}
	return values, nil
}
