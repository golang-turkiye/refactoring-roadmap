package postgres

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewUserRepository(db *gorm.DB, logger *logrus.Logger) (*UserRepository, error) {
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		logger.Error("Error while auto migrating user table")
		return nil, err
	}
	return &UserRepository{db: db, logger: logger}, nil
}

// GetUserByID returns a user by its ID.
func (r *UserRepository) GetUserByID(userID uint) (*domain.User, error) {
	return nil, nil
}

// GetUserByEmail returns a user by its email.
func (r *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	return nil, nil
}
