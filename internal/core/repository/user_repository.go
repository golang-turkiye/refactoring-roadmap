package repository

import "github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"

type UserRepository interface {
	GetUserByID(userID uint) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}
