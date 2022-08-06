package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Golang-Turkiye/refactoring-roadmap/config"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func successfulResponse(w http.ResponseWriter, schema interface{}) {
	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(schema)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(response)
}

func badRequestResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	response, err := json.Marshal(err)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(response)
}

func notFoundResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	response, err := json.Marshal(err)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(response)
}

func notAuthorizedResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	response, err := json.Marshal(err)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(response)
}

func internalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	response, err := json.Marshal(err)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(response)
}

func getUserIDByToken(r *http.Request) (uint, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return 0, errors.New("token not found")
	}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecret), nil
	})
	if err != nil {
		return 0, errors.New("token not valid")
	}
	return uint(claims["user_id"].(float64)), nil
}

func getUserIDByTokenOrBadRequest(r *http.Request) (uint, error) {
	userID, err := getUserIDByToken(r)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
