package service

import "github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"

type UserService interface {
	Login(*domain.User) (string, error)
	GetUser(userID uint) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}
