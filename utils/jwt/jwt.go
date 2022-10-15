package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/arifintahu/go-rest-api/entities"
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
	RoleId uint16 `json:"role_id"`
	RoleSlug string `json:"role_slug"`
	jwt.StandardClaims
}

func GenerateJWT(user *entities.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		RoleId: user.RoleID,
		RoleSlug: user.Role.Slug,
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

func VerifyJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return &Claims{}, err
	}

	if !token.Valid {
		return &Claims{}, ErrTokenInvalid
	}

	return claims, nil
}
