package userservice

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/helpers/utils/authentication"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/urlshorter/domain"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/urlshorter/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

// New creates a new userService.
func New(userRepository repository.UserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

// Login is used to login a user.
func (s *userService) Login(user *domain.User) (string, error) {
	token, err := authentication.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

// GetUser is used to get a user by id.
func (s *userService) GetUser(userID uint) (*domain.User, error) {
	return s.userRepository.GetUserByID(userID)
}

// GetUserByEmail is used to get a user by email.
func (s *userService) GetUserByEmail(email string) (*domain.User, error) {
	return s.userRepository.GetUserByEmail(email)
}
