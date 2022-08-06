package repository

import "github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"

// UserRepository is the repository for users.
type UserRepository interface {
	GetUserByID(userID uint) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}
