package authentication

import (
	"errors"
	jwt "github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

type CustomClaims struct {
	Email string
	jwt.RegisteredClaims
	ExpireAt time.Time
}

func (c *CustomClaims) Valid() error {
	if c.ExpireAt.Unix() < time.Now().Unix() {
		return errors.New("Expired Token")
	}
	if c.Email == "" {
		return errors.New("invalid token")
	}
	return nil
}

func GenerateToken(email string) (string, error) {
	claims := CustomClaims{
		Email:    email,
		ExpireAt: time.Now().Add(24 * time.Hour),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	return token.SignedString([]byte("secret"))
}
func GetEmailByToken(token string) (string, error) {
	arr := strings.Split(token, " ")
	if len(arr) <= 1 {
		return "", errors.New("invalid token")
	}
	token = arr[1]

	claims := CustomClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}
	return claims.Email, nil
}
