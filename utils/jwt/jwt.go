package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	_ "github.com/joho/godotenv/autoload"
)

var secretKey = []byte(os.Getenv("TOKEN_SECRET"))

var (
	ErrTokenParsing = errors.New("Token parsing error")
	ErrTokenInvalid = errors.New("Token invalid")
)

type Claims struct {
	UserId uint64 `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userId uint64) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (uint64, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, ErrTokenInvalid
	}

	return claims.UserId, nil
}
