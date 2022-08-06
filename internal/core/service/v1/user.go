package v1

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (s *UserService) Login(*domain.User) (string, error) {
	return "", nil
}
func (s *UserService) GetUser(userID uint) (*domain.User, error) {
	return nil, nil
}
func (s *UserService) GetUserByEmail(email string) (*domain.User, error) {
	return nil, nil
}
