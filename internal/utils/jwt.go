package utils

import (
	"fmt"
	"time"

	"github.com/brunoeduardodev/simple-ecommerce-go/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user models.User) (string, error) {
	token_lifespan := 24

	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// TODO: Load secret from envs
	return token.SignedString([]byte("TEMPORARY_SECRET"))
}

func ReadToken(tokenString string) (user_id string, err error) {
	// jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {})
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("TEMPORARY_SECRET"), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user_id = claims["user_id"].(string)
		return
	}

	return "", err
}
