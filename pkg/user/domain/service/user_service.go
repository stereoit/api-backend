package service

import (
	"fmt"

	"github.com/stereoit/eventival/pkg/user/domain/repository"
)

// UserService struct for listing dependencies
type UserService struct {
	repo repository.UserRepository
}

// NewUserService returns new instance of UserService
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Duplicated returns error when user with this email already exists
func (s *UserService) Duplicated(email string) error {
	user, err := s.repo.FindByEmail(email)
	if user != nil {
		return fmt.Errorf("%v already exists", email)
	}

	if err != nil {
		return err
	}
	return nil
}
