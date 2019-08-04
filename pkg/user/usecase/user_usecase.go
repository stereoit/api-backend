package usecase

import (
	"github.com/google/uuid"
	"github.com/stereoit/eventival/pkg/user/domain/model"
	"github.com/stereoit/eventival/pkg/user/domain/repository"
	"github.com/stereoit/eventival/pkg/user/domain/service"
)

// UserUsecase defines interface for all user usecases
type UserUsecase interface {
	ListUser() ([]*User, error)
	RegisterUser(email string) error
}

type userUsecase struct {
	repo    repository.UserRepository
	service service.UserService
}

// NewUserUsecase returns new instance of users usecases
func NewUserUsecase(repo repository.UserRepository, service service.UserService) UserUsecase {
	return &userUsecase{
		repo:    repo,
		service: service,
	}
}

func (u *userUsecase) ListUser() ([]*User, error) {
	users, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return toUser(users), nil
}

func (u *userUsecase) RegisterUser(email string) error {
	uid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	if err := u.service.Duplicated(email); err != nil {
		return err
	}
	user := model.NewUser(uid.String(), email)
	if err := u.repo.Save(user); err != nil {
		return err
	}
	return nil
}

// User type defines exported user
type User struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
}

func toUser(users []*model.User) []*User {
	res := make([]*User, len(users))
	for i, user := range users {
		res[i] = &User{
			ID:        user.GetID(),
			Email:     user.GetEmail(),
			FirstName: user.GetFirstName(),
			LastName:  user.GetLastName(),
		}
	}
	return res
}
