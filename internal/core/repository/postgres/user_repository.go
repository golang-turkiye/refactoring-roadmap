package postgres

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// GetUserByID returns a user by its ID.
func (r *UserRepository) GetUserByID(userID uint) (*domain.User, error) {
	return nil, nil
}

// GetUserByEmail returns a user by its email.
func (r *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	return nil, nil
}
