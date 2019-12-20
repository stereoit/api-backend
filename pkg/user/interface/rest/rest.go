package rest

import (
	"fmt"
	"net/http"

	"github.com/stereoit/eventival/pkg/user/registry"

	"github.com/stereoit/eventival/pkg/user/usecase"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
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
	r.Get("/", s.listAllUsers)
	r.Post("/", s.registerUser)

	return r
}

func (s *userService) listAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.userUsecase.ListAllUsers()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	if err := render.RenderList(w, r, NewUsersListResponse(toUserList(users))); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (s *userService) registerUser(w http.ResponseWriter, r *http.Request) {
	userRequest := &UserRequest{}

	if err := render.Bind(r, userRequest); err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}

	createdUser, err := s.userUsecase.RegisterUser(userRequest.Email)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, ErrRender(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewUserResponse(toUser(createdUser)))
}
