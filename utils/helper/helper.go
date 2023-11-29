package helper

import (
	"checkout-api/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateUUID() string {
	return uuid.New().String()
}

/*
return example

	{
		"username": "fulan",
		"user_id": "1234",
		"session_id": "1234",
		"exp": 12345,
	}
*/
func GenerateJwtToken(username string, user_id string, session_id string, secret_key string, exp int) (string, error) {
	secretKey := []byte(secret_key)

	claims := jwt.MapClaims{
		"user_id":    user_id,
		"username":   username,
		"exp":        time.Now().Add(time.Hour * time.Duration(exp)).Unix(),
		"session_id": session_id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	var JWT_SIGNATURE_KEY = []byte(config.Envs.JWT_SECRET)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}
