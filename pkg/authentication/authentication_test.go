package authentication_test

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/pkg/authentication"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	mail := "test@mail.test"
	token, err := authentication.GenerateToken(mail)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
	claims := authentication.CustomClaims{}
	_, err = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	assert.Nil(t, err)
	assert.Equal(t, claims.Email, mail)
}

func TestGetEmailByToken(t *testing.T) {
	mail := "test@mail.test"
	claims := authentication.CustomClaims{
		Email: mail,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	assert.Nil(t, err)
	assert.NotEmpty(t, tokenString)
	email, err := authentication.GetEmailByToken(tokenString)
	assert.Nil(t, err)
	assert.Equal(t, email, mail)
}
