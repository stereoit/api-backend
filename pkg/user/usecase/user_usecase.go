package usecase

import (
	"github.com/google/uuid"
	"github.com/stereoit/eventival/pkg/user/domain"
	"github.com/stereoit/eventival/pkg/user/domain/model"
	"github.com/stereoit/eventival/pkg/user/domain/repository"
	"github.com/stereoit/eventival/pkg/user/domain/service"
)

// UserUsecase defines interface for all user usecases
type UserUsecase interface {
	ListAllUsers() ([]*User, error)
	RegisterUser(email string) (string, error)
	FindByID(id string) (*User, error)
	UpdateUser(*User) error
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

func (u *userUsecase) ListAllUsers() ([]*User, error) {
	users, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return toUsecaseUserList(users), nil
}

func (u *userUsecase) FindByID(id string) (*User, error) {
	user, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return toUsecaseUser(user), nil
}

func (u *userUsecase) RegisterUser(email string) (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	if err := u.service.Duplicated(email); err != nil {
		return "", &domain.DuplicateError{}
	}
	user := model.NewUser(uid.String(), email)
	if err := u.repo.Save(user); err != nil {
		return "", err
	}
	return user.GetID(), nil
}

// UpdateUser updates the user entity with validation
func (u *userUsecase) UpdateUser(user *User) error {
	entity := toUser(user)

	if err := entity.Validate(); err != nil {
		// log.Printf("Error: %v\n", err)
		return err
	}

	if err := u.repo.Update(toUser(user)); err != nil {
		// log.Printf("Error: %v\n", err)
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

func toUser(user *User) *model.User {
	entity := model.NewUser(user.ID, user.Email)
	entity.SetFirstName(user.FirstName)
	entity.SetLastName(user.LastName)

	return entity
}

func toUsecaseUser(user *model.User) *User {
	return &User{
		ID:        user.GetID(),
		Email:     user.GetEmail(),
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
	}
}

func toUsecaseUserList(users []*model.User) []*User {
	res := make([]*User, len(users))
	for i, user := range users {
		res[i] = toUsecaseUser(user)
	}
	return res
}
