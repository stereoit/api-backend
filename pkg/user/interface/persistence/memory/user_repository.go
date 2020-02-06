package memory

import (
	"errors"
	"sync"

	"github.com/stereoit/eventival/pkg/user/domain/model"
	"github.com/stereoit/eventival/pkg/user/domain/repository"
)

type userRepository struct {
	mu    *sync.Mutex
	users map[string]*User
}

// NewUserRepository returns new instance of memory based repository
func NewUserRepository() repository.UserRepository {
	return &userRepository{
		mu:    &sync.Mutex{},
		users: map[string]*User{},
	}
}

func (r *userRepository) FindAll() ([]*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	users := make([]*model.User, len(r.users))
	i := 0
	for _, user := range r.users {
		users[i] = model.NewUser(user.ID, user.Email)
		i++
	}

	return users, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.Email == email {
			return model.NewUser(user.ID, user.Email), nil
		}
	}
	return nil, nil
}

func (r *userRepository) FindByID(id string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.ID == id {
			return model.NewUser(user.ID, user.Email), nil
		}
	}
	return nil, nil
}

func (r *userRepository) Save(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.GetID()] = &User{
		ID:        user.GetID(),
		Email:     user.GetEmail(),
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
	}
	return nil
}

func (r *userRepository) Update(user *model.User) error {
	return r.Save(user)
}

func (r *userRepository) Delete(userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[userID]; !ok {
		return errors.New("User does not exists")
	}

	delete(r.users, userID)
	return nil
}

// User struct maps the model.User
type User struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
}
