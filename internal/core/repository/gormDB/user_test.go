package gormDB

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"
	"github.com/Golang-Turkiye/refactoring-roadmap/pkg/database/localDB"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestNewUserRepository(t *testing.T) {
	db, err := localDB.Connection("test_new_user_repository.db")
	assert.Nil(t, err)
	userRepo, err := NewUserRepository(db, logrus.New())
	assert.Nil(t, err)
	assert.NotNil(t, userRepo)
	assert.Nil(t, os.Remove("test_new_user_repository.db"))
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	db, err := localDB.Connection("test_get_user_by_email.db")
	assert.Nil(t, err)
	userRepo, err := NewUserRepository(db, logrus.New())
	assert.Nil(t, err)
	users := []domain.User{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Email:    "not+correct@gmail.com",
			Password: "123456",
		},
		{
			Model: gorm.Model{
				ID: 99,
			},
			Email:    "alameddinc+test@gmail.com",
			Password: "123456",
		},
	}
	for _, user := range users {
		assert.Nil(t, db.Create(&user).Error)
	}
	t.Run("Success", func(t *testing.T) {
		user, err := userRepo.GetUserByEmail(users[1].Email)
		assert.Nil(t, err)
		assert.Equal(t, user.Email, "alameddinc+test@gmail.com")
		assert.Equal(t, user.Password, "123456")
	})
	assert.Nil(t, os.Remove("test_get_user_by_email.db"))
}

func TestUserRepository_GetUserByID(t *testing.T) {
	db, err := localDB.Connection("test_get_user_by_id.db")
	assert.Nil(t, err)
	userRepo, err := NewUserRepository(db, logrus.New())
	assert.Nil(t, err)
	users := []domain.User{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Email:    "not+correct@gmail.com",
			Password: "123456",
		},
		{
			Model: gorm.Model{
				ID: 99,
			},
			Email:    "alameddinc+test@gmail.com",
			Password: "123456",
		},
	}
	for _, user := range users {
		assert.Nil(t, db.Create(&user).Error)
	}
	t.Run("Success", func(t *testing.T) {
		user, err := userRepo.GetUserByID(users[1].ID)
		assert.Nil(t, err)
		assert.Equal(t, user.Email, "alameddinc+test@gmail.com")
		assert.Equal(t, user.Password, "123456")
	})
	assert.Nil(t, os.Remove("test_get_user_by_id.db"))
}
