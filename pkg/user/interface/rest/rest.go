package rest

import (
	"net/http"

	"github.com/stereoit/eventival/pkg/user/registry"

	"github.com/stereoit/eventival/pkg/user/usecase"

	"github.com/go-chi/chi"
)

// Apply adds this resource to the given router
func Apply(path string, router chi.Router, ctn *registry.Container) {
	userService := NewUserService(ctn.Resolve("user-usecase").(usecase.UserUsecase))
	router.Mount(path, userService.Routes())
}

type userService struct {
	userUsecase usecase.UserUsecase
}

// UserService interfaces describes Application Service for Users
type UserService interface {
	Routes() chi.Router
}

// NewUserService returns instance of UserService
func NewUserService(userUsecase usecase.UserUsecase) UserService {
	return &userService{
		userUsecase: userUsecase,
	}
}

func (s *userService) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", s.list)

	return r
}

func (s *userService) list(w http.ResponseWriter, r *http.Request) {
	users, err := s.userUsecase.ListUser()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	for _, user := range users {
		w.Write([]byte(user.Email))
	}
}
