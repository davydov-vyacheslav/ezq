package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID string   `json:"sub"`
	Email  string   `json:"email"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID, email string, roles []string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "ezq-backend",
			Audience:  []string{"ezq-frontend"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(jwtSecret)
}

func ParseAccessToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)

	if err != nil || !token.Valid {
		return nil, err
	}

	return token.Claims.(*Claims), nil
}
