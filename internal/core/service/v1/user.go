package v1

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/repository"
	"github.com/Golang-Turkiye/refactoring-roadmap/pkg/authentication"
)

type UserService struct {
	userRepository repository.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// Login is used to login a user.
func (s *UserService) Login(user *domain.User) (string, error) {
	token, err := authentication.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

// GetUser is used to get a user by id.
func (s *UserService) GetUser(userID uint) (*domain.User, error) {
	return s.userRepository.GetUserByID(userID)
}

// GetUserByEmail is used to get a user by email.
func (s *UserService) GetUserByEmail(email string) (*domain.User, error) {
	return s.userRepository.GetUserByEmail(email)
}
