package v1_test

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/repository/mocks"
	v1 "github.com/Golang-Turkiye/refactoring-roadmap/internal/core/service/v1"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestNewUserService(t *testing.T) {
	userRepo := new(mocks.MockUserRepository)
	userService := v1.NewUserService(userRepo)
	assert.NotNil(t, userService)
}

func TestUserService_GetUser(t *testing.T) {
	userRepo := new(mocks.MockUserRepository)
	userService := v1.NewUserService(userRepo)
	assert.NotNil(t, userService)
	testCases := []struct {
		name     string
		user     *domain.User
		expected error
	}{
		{
			name: "Success",
			user: &domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Email:    "alameddinc+test@gmail.com",
				Password: "123456",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo.On("GetUserByID", tc.user.ID).Return(tc.user, tc.expected)
			user, err := userService.GetUser(tc.user.ID)
			assert.Equal(t, err, tc.expected)
			assert.Equal(t, user, tc.user)
		})
	}
}

func TestUserService_GetUserByEmail(t *testing.T) {
	userRepo := new(mocks.MockUserRepository)
	userService := v1.NewUserService(userRepo)
	assert.NotNil(t, userService)
	testCases := []struct {
		name     string
		user     *domain.User
		expected error
	}{
		{
			name: "Success",
			user: &domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Email:    "alameddinc+test@gmail.com",
				Password: "123456",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo.On("GetUserByEmail", tc.user.Email).Return(tc.user, tc.expected)
			user, err := userService.GetUserByEmail(tc.user.Email)
			assert.Equal(t, err, tc.expected)
			assert.Equal(t, user, tc.user)
		})
	}
}

func TestUserService_Login(t *testing.T) {
	userRepo := new(mocks.MockUserRepository)
	userService := v1.NewUserService(userRepo)
	assert.NotNil(t, userService)
	testCases := []struct {
		name     string
		user     *domain.User
		expected error
	}{
		{
			name: "Success",
			user: &domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Email:    "alameddinc+test@gmail.com",
				Password: "123456",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo.On("GetUserByEmail", tc.user.Email).Return(tc.user, tc.expected)
			token, err := userService.Login(tc.user)
			assert.Equal(t, err, tc.expected)
			assert.NotEmpty(t, token)
		})
	}
}
