package helpers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go-marketplace/models"
	"time"
)

var sigInKey = []byte("mySecret")

func CreateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username":  user.Email,
			"name":      user.Name,
			"expiredAt": time.Now().Add(time.Hour * 24).Unix(),
			"issuedAt":  time.Now(),
			"notBefore": time.Now(),
		})

	tokenString, err := token.SignedString(sigInKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return sigInKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
