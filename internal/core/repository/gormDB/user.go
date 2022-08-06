package gormDB

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
	tx := r.db.Begin()
	user := &domain.User{}
	result := tx.First(user, userID)
	if result.Error != nil || result.RowsAffected == 0 {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return user, nil
}

// GetUserByEmail returns a user by its email.
func (r *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	tx := r.db.Begin()
	user := &domain.User{}
	result := tx.Where("email = ?", email).First(user)
	if result.Error != nil || result.RowsAffected == 0 {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return user, nil
}
