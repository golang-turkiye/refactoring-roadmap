package controller

import (
	"errors"
	"github.com/Golang-Turkiye/refactoring-roadmap/config"
	"github.com/Golang-Turkiye/refactoring-roadmap/db/mysql"
	"github.com/Golang-Turkiye/refactoring-roadmap/model"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func PostLogin(w http.ResponseWriter, r *http.Request) {
	conn, err := mysql.Connection()
	if err != nil {
		internalServerError(w, err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		badRequestResponse(w, errors.New("username or password is empty"))
		return
	}
	user, err := model.GetUserByUsernamePassword(conn, username, password)
	if err != nil {
		notFoundResponse(w, err)
		return
	}
	if user.Password != password {
		notAuthorizedResponse(w, errors.New("password is wrong"))
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		badRequestResponse(w, err)
		return
	}
	successfulResponse(w, map[string]string{"token": tokenString})
}
