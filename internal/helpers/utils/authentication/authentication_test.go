package authentication_test

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/helpers/utils/authentication"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	mail := "test@mail.test"
	token, err := authentication.GenerateToken(mail)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
	claims := authentication.CustomClaims{Email: mail}
	_, err = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	assert.Nil(t, err)
	assert.Equal(t, claims.Email, mail)
}

func TestGetEmailByToken(t *testing.T) {
	mail := "test@mail.test"
	claims := authentication.CustomClaims{Email: mail, ExpireAt: time.Now().Add(10 * time.Minute)}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, err := token.SignedString([]byte("secret"))
	assert.Nil(t, err)
	assert.NotEmpty(t, tokenString)
	email, err := authentication.GetEmailByToken("bearer " + tokenString)
	assert.Nil(t, err)
	assert.Equal(t, email, mail)
}
