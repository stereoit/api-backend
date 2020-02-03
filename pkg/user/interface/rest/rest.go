package rest

import (
	"net/http"

	"github.com/sarulabs/di"

	"github.com/stereoit/eventival/pkg/user/usecase"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Apply adds this resource to the given router
func Apply(path string, router chi.Router, ctn di.Container) {
	userService := New(ctn.Get("user-usecase").(usecase.UserUsecase))
	router.Mount(path, userService.Routes())
}

type userService struct {
	userUsecase usecase.UserUsecase
}

// UserService interfaces describes Application Service for Users
type UserService interface {
	Routes() chi.Router
}

// New returns instance of UserService
func New(userUsecase usecase.UserUsecase) UserService {
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
	userRequest := &RegisterUserRequest{}

	if err := render.Bind(r, userRequest); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	newUserID, err := s.userUsecase.RegisterUser(userRequest.Email)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewRegisterUserResponse(newUserID))
}
