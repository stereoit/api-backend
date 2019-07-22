package memory

import (
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
		users[i] = model.NewUser(user.ID, user.Email, user.FirstName, user.LastName)
		i++
	}

	return users, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.Email == email {
			return model.NewUser(user.ID, user.Email, user.FirstName, user.LastName), nil
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

// User struct maps the model.User
type User struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
}
