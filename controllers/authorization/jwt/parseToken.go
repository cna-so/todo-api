package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func ParseToken(jwtToken string) (*jwt.Token, error) {
	token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})
	return token, nil
}
