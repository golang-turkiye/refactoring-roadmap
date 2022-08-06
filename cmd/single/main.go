package main

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/handler/apiV1"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/repository/gormDB"
	v1 "github.com/Golang-Turkiye/refactoring-roadmap/internal/core/service/v1"
	"github.com/Golang-Turkiye/refactoring-roadmap/pkg/database/localDB"
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
	linkService := v1.NewLinkService(linkRepo)
	userService := v1.NewUserService(userRepo)
	// Handlers Generated
	handler := apiV1.NewShortURLHander(linkService, userService, router, logger)
	handler.Run()
}
