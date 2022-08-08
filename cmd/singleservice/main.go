package main

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/urlshorter/handler/api/v1/shorturlhandler"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/urlshorter/repository/gormDB"
	linkservice "github.com/Golang-Turkiye/refactoring-roadmap/internal/urlshorter/service/v1/link"
	userservice "github.com/Golang-Turkiye/refactoring-roadmap/internal/urlshorter/service/v1/user"
	"github.com/Golang-Turkiye/refactoring-roadmap/src/database/localDB"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := localDB.Connection("mock_db.db")
	if err != nil {
		panic(err)
	}

	// Router & Logger Generated
	router := mux.NewRouter()
	logger := logrus.New()
	// Repositories Generated
	userRepo, err := gormDB.NewUserRepository(db, logger)
	if err != nil {
		panic(err)
	}
	linkRepo, err := gormDB.NewLinkRepository(db, logger)
	if err != nil {
		panic(err)
	}
	// Services Generated
	linkService := linkservice.New(linkRepo)
	userService := userservice.New(userRepo)
	// Handlers Generated
	handler := shorturlhandler.NewShortURLHander(linkService, userService, router, logger)
	handler.Run()
}
