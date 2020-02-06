package repository

import "github.com/stereoit/eventival/pkg/user/domain/model"

// UserRepository describe interface
type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByID(id string) (*model.User, error)
	Save(user *model.User) error
	Update(user *model.User) error
	Delete(id string) error
}
